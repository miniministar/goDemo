package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

func initGormDB() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	//设置全局表名禁用复数
	db.SingularTable(true)
	return db
}

func TestGormConnect(t *testing.T) {
	db := initGormDB()
	if db == nil {
		return
	}
	db.AutoMigrate(&user{})
}

func TestGormInsert(t *testing.T) {
	db := initGormDB()
	if db == nil {
		return
	}
	u := user{Name: "你是谁", Age: 1}
	//这里使用了Table()函数，如果你没有指定全局表名禁用复数，或者是表名跟结构体名不一样的时候
	//你可以自己在sql中指定表名。这里是示例，本例中这个函数可以去除。
	db.Table("user").Create(&u)
	//tempResult := db.Table("user").Create(user)
	//fmt.Printf("%v\n", tempResult)
}

func TestGormUpdate(t *testing.T) {
	db := initGormDB()
	if db == nil {
		return
	}
	u := user{Id: 1, Name: "小明"}
	db.Model(&u).Update(u)

	db.Model(&user{}).Where("id = ?", 100).Update("name", "小白")
}

func TestGormDelete(t *testing.T) {
	db := initGormDB()
	if db == nil {
		return
	}
	u := user{Id: 4}
	i := db.Delete(&u).RowsAffected
	fmt.Printf("delete affect rows %d\n", i)
}

func TestGormQuery(t *testing.T) {
	db := initGormDB()
	if db == nil {
		return
	}
	var u user
	db.Find(&u, "id=?", 100)
	printUser(&u)
	var u2 []user
	//db.Where("id in( ?)", [] int{100, 101}).Find(&u2)
	//for i := 0; i < len(u2); i++ {
	//	printUser(&u2[i])
	//}

	db.Select("id, age").Where(map[string]interface{}{"age": 100, "name": "历史"}).Find(&u2)
	for i := 0; i < len(u2); i++ {
		printUser(&u2[i])
	}
}
