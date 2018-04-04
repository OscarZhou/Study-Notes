package configs

import (
	"log"

	config "github.com/zpatrick/go-config"
)

type localConfig struct {
	Port                    string
	Production              bool
	DbConn                  string
	DbEngin                 string
	DropTables              bool
	AutoMigrate             bool
	TRKDApplicationID       string
	TRKDUserName            string
	TRKDPassword            string
	ServiceTokenURL         string
	HeaderHost              string
	ForexRateURL            string
	ForexRateRequestKeyName string
	ForexRateFields         string
	ForexRateScope          string
	DefaultBasisPoint       int
	DefaultPercentage       int
}

var LocalConfig localConfig

func init() {
	localIniFile := config.NewINIFile("configs/local.ini")
	local := config.NewConfig([]config.Provider{localIniFile})
	if err := local.Load(); err != nil {
		panic(err)
	}

	var err error

	LocalConfig.Production, err = local.Bool("local.production")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	LocalConfig.DbConn, err = local.String("local.db_conn")
	if err != nil {
		panic(err)
	}
	LocalConfig.DbEngin, err = local.String("local.db_engin")
	if err != nil {
		panic(err)
	}

	LocalConfig.DropTables, err = local.Bool("local.drop_tables")
	if err != nil {
		panic(err)
	}

	LocalConfig.AutoMigrate, err = local.Bool("local.auto_migrate")
	if err != nil {
		panic(err)
	}

	LocalConfig.TRKDApplicationID, err = local.String("local.trkd_application_id")
	if err != nil {
		panic(err)
	}

	LocalConfig.TRKDUserName, err = local.String("local.trkd_username")
	if err != nil {
		panic(err)
	}

	LocalConfig.TRKDPassword, err = local.String("local.trkd_password")
	if err != nil {
		panic(err)
	}

	LocalConfig.ServiceTokenURL, err = local.String("local.service_token_url")
	if err != nil {
		panic(err)
	}

	LocalConfig.HeaderHost, err = local.String("local.header_host")
	if err != nil {
		panic(err)
	}

	LocalConfig.ForexRateURL, err = local.String("local.forex_rate_url")
	if err != nil {
		panic(err)
	}

	LocalConfig.ForexRateRequestKeyName, err = local.String("local.forex_rate_request_key_name")
	if err != nil {
		panic(err)
	}

	LocalConfig.ForexRateFields, err = local.String("local.forex_rate_fields")
	if err != nil {
		panic(err)
	}

	LocalConfig.ForexRateScope, err = local.String("local.forex_rate_scope")
	if err != nil {
		panic(err)
	}

}
