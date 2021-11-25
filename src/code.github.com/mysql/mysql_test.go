package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var db *sql.DB

func initDB() (*sql.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8&parseTime=True"
	// 不会校验账号密码是否正确
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	fmt.Printf("db %T\n", db)
	if err != nil {
		return db, err
	}
	return db, err
}

func TestConnect(t *testing.T) {
	_, err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
	}
}

func TestInsert(t *testing.T) {
	db, err := initDB()
	if err != nil {
		fmt.Printf("connect db is failed, error %v", err)
		return
	}
	sqlStr := "insert into user(name, age) values(?,?)"
	for i := 0; i < 100; i++ {
		item := i + 1
		db.Exec(sqlStr, fmt.Sprintf("张三%d", item), 38+item)
	}
	result, err := db.Exec(sqlStr, "张三", 38)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	fmt.Printf("id %d\n", id)

}

func TestQuery(t *testing.T) {
	db, err := initDB()
	fmt.Printf("%T %v\n", db, err)
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	row := db.QueryRow(sqlStr, 1)
	if row == nil {
		fmt.Printf("query id %d row is empty \n", 1)
		return
	}
	err = row.Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err :%v\n", err)
		return
	}
	if &u != nil {
		fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
	}
	defer db.Close()
}

func TestUpdate(t *testing.T) {
	db, err := initDB()
	fmt.Printf("%T %v\n", db, err)
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	sqlStr := "update user set age = ? where id = ?"
	result, err := db.Exec(sqlStr, 25, 1)
	if err != nil {
		fmt.Printf("update is failed , error %v\n", err)
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("update affect rows %d\n", rows)
}

func TestDelete(t *testing.T) {
	db, err := initDB()
	fmt.Printf("%T %v\n", db, err)
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	sqlStr := "delete from user where id = ?"
	result, err := db.Exec(sqlStr, 2)
	if err != nil {
		fmt.Printf("delete user failed, err:%v\n", err)
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("delete affect rows %d\n", rows)
}

func TestPrepare(t *testing.T) {
	db, err := initDB()
	fmt.Printf("%T %v\n", db, err)
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
	sqlStr := "select id ,name,age from user where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed , err %v\n", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(100)
	if err != nil {
		fmt.Printf("query failed, err %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed , err %v \n", err)
			return
		}
		printUser(&u)
	}

}

func printUser(u *user) {
	if u != nil {
		fmt.Printf("id %d name %s age %d\n", u.Id, u.Name, u.Age)
	}
}
