package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func initDb() (db2 *sqlx.DB) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True"
	db2, error := sqlx.Connect("mysql", dsn)
	if error != nil {
		fmt.Printf("connect db failed, err:%v\n", error)
		return db2
	}
	db2.SetMaxOpenConns(20)
	db2.SetMaxIdleConns(10)
	return db2
}

func TestSqlxQuery(t *testing.T) {
	db2 := initDb()
	if db2 == nil {
		return
	}
	sql := "select id ,name, age from user where id = ?"
	var u user
	err := db2.Get(&u, sql, 1)
	if err != nil {
		fmt.Printf("get failed, err : %v\n", err)
		return
	}
	printUser(&u)
}

func TestSqlxSelect(t *testing.T) {
	db2 := initDb()
	if db2 == nil {
		return
	}
	sql := "select id ,name, age from user where id > ?"
	var users []user
	err := db2.Select(&users, sql, 100)
	if err != nil {
		fmt.Printf("get failed, err : %v\n", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
	//printUser(&u)
}

func TestSqlxInsert(t *testing.T) {
	db2 := initDb()
	if db2 == nil {
		return
	}
	sql := "insert into user(name, age) values (?, ?)"
	execResult, err := db2.Exec(sql, "历史", 100)
	if err != nil {
		fmt.Printf("insert failed, err : %v\n", err)
		return
	}
	result, _ := execResult.LastInsertId()
	fmt.Printf("result:%#v\n", result)
}

func TestSqlxUpdate(t *testing.T) {
	db2 := initDb()
	if db2 == nil {
		return
	}
	sql := "update user set age = ? where id = ?"
	execResult, err := db2.Exec(sql, 0, 100)
	if err != nil {
		fmt.Printf("update failed, err : %v\n", err)
		return
	}
	result, _ := execResult.RowsAffected()
	fmt.Printf("result:%#v\n", result)
}

func TestSqlxDelete(t *testing.T) {
	db2 := initDb()
	if db2 == nil {
		return
	}
	sql := "delete from user where id = ?"
	execResult, err := db2.Exec(sql, 3)
	if err != nil {
		fmt.Printf("delete failed, err : %v\n", err)
		return
	}
	result, _ := execResult.RowsAffected()
	fmt.Printf("result:%#v\n", result)
}
