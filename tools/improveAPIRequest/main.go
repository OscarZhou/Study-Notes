package main

import (
	"Study-Notes/tools/improveAPIRequest/configs"
	"Study-Notes/tools/improveAPIRequest/controllers"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer db.Close()
	if configs.LocalConfig.Production {
		gin.SetMode(gin.ReleaseMode)
		db.LogMode(false)
	} else {
		gin.SetMode(gin.DebugMode)
		db.LogMode(true)
	}

	start := time.Now()
	statusCode, err := controllers.UpdateExchangeRatesPeriodically(db)
	if err != nil {
		log.Printf("method name: utils.UpdateExchangeRatesPeriodically, statusCode = %v, error: %v", statusCode, err)
		return
	}

	elapsed := time.Since(start)
	log.Printf("updating exchange rate took %s\n", elapsed)

}
