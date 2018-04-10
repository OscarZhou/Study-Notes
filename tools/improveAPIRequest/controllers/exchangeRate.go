package controllers

import (
	"Study-Notes/tools/improveAPIRequest/configs"
	"Study-Notes/tools/improveAPIRequest/models"
	"Study-Notes/tools/improveAPIRequest/utils"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	currencyList = make(map[string]bool)
	ricList      = make(map[string]bool)
)

func UpdateExchangeRatesPeriodically(db *gorm.DB) (int, error) {
	start := time.Now()
	var (
		currencyRelations []models.CurrencyRelation
		fxResponse        []models.TRKDFXResponse
		exchangeRates     []models.ExchangeRate
		ricNames          []string
	)

	// sql := "SELECT DISTINCT(trkdric) FROM " + utils.GetDBTableNameByModel(db, &models.CurrencyRelation{})
	// if err := db.Raw(sql).Scan(&currencyRelations).Error; err != nil {
	// 	return 0, err
	// }

	if err := db.Find(&currencyRelations).Error; err != nil {
		return 0, err
	}

	for _, v := range currencyRelations {
		currencyList[v.TRKDCurrencyCode] = true

		if _, ok := ricList[v.TRKDRIC]; !ok {
			ricList[v.TRKDRIC] = true
		}
	}

	keys := reflect.ValueOf(ricList).MapKeys()
	for i := 0; i < len(keys); i++ {
		ricNames = append(ricNames, keys[i].String())
	}

	ch := make(chan models.TRKDFXResponse)
	var indexStart = 0
	var indexEnd = 0
	numRic := 2
	for i := 0; i < len(ricNames)/numRic; i++ {
		var rn []string
		if i != len(ricNames)/numRic-1 {
			indexStart = indexEnd
			indexEnd = indexEnd + numRic
			rn = append(rn, ricNames[indexStart:indexEnd]...)
		} else {
			rn = append(rn, ricNames[indexEnd:]...)
		}
		fmt.Println("go routine = ", rn)
		go func() {
			_, err := utils.RequestServiceTokenFromTRKD(rn, ch)
			if err != nil {
				log.Println("go routine return error ")
			}
		}()
	}

	for i := 0; i < len(ricNames)/numRic; i++ {
		fxResponse = append(fxResponse, <-ch)
	}

	for _, v := range fxResponse {
		for _, items := range v.RetrieveItemResponse.ItemResponses[0].Items {
			ricName := items.RequestKey.Name[0:3]
			for _, childItems := range items.ChildItems {
				var (
					from                    string
					to                      string
					conversion              int64
					isAbleCheckExchangeRate = false
				)
				for _, field := range childItems.Fields.Field {
					if field.Name == "GV2_TEXT" {
						from = field.Utf8String[0:3]
						to = field.Utf8String[3:]
						if from == ricName {
							if _, ok := currencyList[to]; !ok {
								break
							}
							isAbleCheckExchangeRate = true
						}
						if to == ricName {
							if _, ok := currencyList[from]; !ok {
								break
							}
							isAbleCheckExchangeRate = true
						}
						continue
					}

					if field.Name == "CF_CLOSE" && isAbleCheckExchangeRate {
						percentage := float64(1.0 + float64(configs.LocalConfig.DefaultBasisPoint)*0.0001)
						conversion = int64(field.Double * percentage * 1000000)
						inverseConversion := int64((1.0 / field.Double) * percentage * 1000000)
						exchangeRate := models.ExchangeRate{
							CurrencyCodeFrom:       from,
							CurrencyCodeTo:         to,
							ConversionValue:        conversion,
							InverseConversionValue: inverseConversion,
						}
						exchangeRates = append(exchangeRates, exchangeRate)
						break
					}
				}
			}
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("request API and iterates all data took %s\n", elapsed)
	start = time.Now()
	insertExchangeRateSQL := `INSERT INTO ` + utils.GetDBTableNameByModel(db, &models.ExchangeRate{}) +
		` (currency_code_from, currency_code_to, conversion_value, inverse_conversion_value) VALUES `
	for _, v := range exchangeRates {
		erData := `('` + v.CurrencyCodeFrom +
			`','` + v.CurrencyCodeTo +
			`',` + strconv.FormatInt(v.ConversionValue, 10) +
			`,` + strconv.FormatInt(v.InverseConversionValue, 10) + `),`
		insertExchangeRateSQL += erData
	}
	if err := db.Exec(insertExchangeRateSQL[0 : len(insertExchangeRateSQL)-1]).Error; err != nil {
		return 500, err
	}
	elapsed = time.Since(start)
	fmt.Printf("insert action took %s\n", elapsed)
	return 200, nil
}
