package clickhouse

import (
	"log"
	"os"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var (
		err   error
		dbDSN = "tcp://10.0.2.34:9000?database=default&username=&password=&read_timeout=10&write_timeout=20"
	)

	if DB, err = gorm.Open(clickhouse.Open(dbDSN), &gorm.Config{}); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}

	if os.Getenv("DEBUG") == "true" {
		DB = DB.Debug()
	}
}

