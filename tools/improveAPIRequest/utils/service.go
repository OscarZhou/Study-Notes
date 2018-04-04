package utils

import (
	"Study-Notes/tools/improveAPIRequest/configs"
	"Study-Notes/tools/improveAPIRequest/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func RequestServiceTokenFromTRKD(ricNames []string, ch chan<- models.TRKDFXResponse) (int, error) {
	var (
		requestKeys []models.RequestKey
		fxResponse  models.TRKDFXResponse
	)
	start := time.Now()
	url := configs.LocalConfig.ServiceTokenURL
	header := map[string]string{
		"Content-Type":   "application/json",
		"Host":           configs.LocalConfig.HeaderHost,
		"Accept_Charset": "utf-8",
	}

	serviceTokenRequest := models.ServiceTokenRequest{
		ApplicationID: configs.LocalConfig.TRKDApplicationID,
		Username:      configs.LocalConfig.TRKDUserName,
		Password:      configs.LocalConfig.TRKDPassword,
	}

	tokenRequest := models.TRKDTokenRequest{
		ServiceTokenRequest: serviceTokenRequest,
	}

	tokenRequestString, err := ConvertJSON2String(tokenRequest)
	if err != nil {
		return 500, err
	}

	body := strings.NewReader(tokenRequestString)

	var trkdToken models.TRKDTokenResponse
	statusCode, err := Request("POST", url, header, body, &trkdToken, 3)
	if err != nil {
		return statusCode, err
	}

	url = configs.LocalConfig.ForexRateURL
	header = map[string]string{
		"Content-Type":              "application/json",
		"X-Trkd-Auth-Token":         trkdToken.ServiceTokenResponse.Token,
		"X-Trkd-Auth-ApplicationID": configs.LocalConfig.TRKDApplicationID,
	}

	for _, v := range ricNames {
		requestKeys = append(requestKeys, models.RequestKey{Name: v, NameType: "RIC"})
	}
	FXRequest := models.TRKDFXRequest{
		RetrieveItemRequest: models.RetrieveItemRequest{
			ItemRequests: []models.ItemRequest{
				models.ItemRequest{
					Fields:      configs.LocalConfig.ForexRateFields,
					RequestKeys: requestKeys,
					Scope:       configs.LocalConfig.ForexRateScope,
				},
			},
			TrimResponse:        false,
			IncludeChildItemQoS: false,
		},
	}
	fxRequestString, err := ConvertJSON2String(FXRequest)
	if err != nil {
		return 500, err
	}

	body = strings.NewReader(fxRequestString)
	statusCode, err = Request("POST", url, header, body, &fxResponse, 3)
	if err != nil {
		return statusCode, err
	}
	elapsed := time.Since(start)
	fmt.Printf("request Reuters' API took %s\n", elapsed)
	ch <- fxResponse
	return 200, nil
}

func Request(method, urlStr string, header interface{}, body io.Reader, object interface{}, retry int) (int, error) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req, _ := http.NewRequest(method, urlStr, body)
	req = req.WithContext(c)

	headerValue := reflect.ValueOf(header)
	if headerValue.IsValid() {
		switch headerValue.Kind() {
		case reflect.Map:
			keys := headerValue.MapKeys()
			for _, k := range keys {
				v := headerValue.MapIndex(k)
				// Check if both of the key and the value are all string type
				if k.Type().Kind() == reflect.String && v.Type().Kind() == reflect.String {
					switch strings.ToLower(k.String()) {
					case "host":
						req.Header.Add("Host", v.String())
						break
					case "content-type":
						req.Header.Add("Content-Type", v.String())
						break
					case "accept-charset":
						req.Header.Add("Accept-Charset", v.String())
						break
					case "x-trkd-auth-token":
						req.Header.Add("X-Trkd-Auth-Token", v.String())
						break
					case "x-trkd-auth-applicationid":
						req.Header.Add("X-Trkd-Auth-ApplicationID", v.String())
						break
					}
				} else {
					return 400, errors.New("The data type in request header is not valid ")
				}
			}
			break
		case reflect.String:
			break
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 500, err
	}
	if resp.StatusCode >= 500 {
		if retry < 2 {
			return resp.StatusCode, errors.New(resp.Status)
		}
		return Request(method, urlStr, header, body, object, retry-1)
	}
	if resp.StatusCode >= 400 && resp.StatusCode < 500 || resp.StatusCode < 200 {
		defer resp.Body.Close()
		resBody, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
		if err != nil {
			return 500, err
		}
		stringBody := string(resBody)
		return resp.StatusCode, errors.New(stringBody)
	}
	if resp.StatusCode == 204 {
		return 204, nil
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		return 500, err
	}
	stringBody := string(resBody)
	if stringBody == "" {
		return 500, io.EOF
	}
	err = json.Unmarshal([]byte(resBody), object)
	if err != nil {
		return 500, err
	}

	return resp.StatusCode, nil
}
