package clickhouse

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

func initClickDB() (db *gorm.DB) {
	var (
		err   error
		dbDSN = "tcp://10.0.50.100:9000?database=newdb&read_timeout=10&write_timeout=20"
	)

	if db, err = gorm.Open(clickhouse.Open(dbDSN), &gorm.Config{}); err != nil {
		log.Printf("failed to connect database, got error %v", err)
		os.Exit(1)
	}
	return db
}

type Test struct {
	Id    int
	Money string `xorm: "decimal(18,4)"`
}

type Book struct {
	Id    uint `gorm:"primary_key"`
	Title string
	Price decimal.Decimal `gorm: "type:decimal(20,8);"`
}

func TestDecimal(t *testing.T) {
	db := initClickDB()

	//db.Exec(ddl)
	//tx := db.Begin()
	//tx.Exec(dml, 10.0, 10.1)
	var c Test
	//db.Exec(query).Scan(&c)
	db.Table("test").Find(&c)
	fmt.Printf("%v\n", c.Money)

	//db.AutoMigrate(&Book{})

	//b := Book{
	//	Id:1,
	//	Title:"title",
	//	Price: decimal.NewFromFloat32(111.234),
	//}
	//db.Create(&b)
	var b1 Book
	db.Where("id=?", 1).Find(&b1)
	println("%v%v%v", b1.Id, b1.Price.InexactFloat64(), b1.Title)
}

func TestCK(t *testing.T) {

	connect, err := sqlx.Open("clickhouse", "tcp://10.0.50.100:9000?debug=true")
	if err != nil {
	}

	sql := "select * from newdb.test where id = 0"
	var items []Test

	if err := connect.Select(&items, sql); err != nil {

	}

	println("items%v", items)
}
