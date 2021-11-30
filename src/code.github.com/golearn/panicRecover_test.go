package main

import (
	"fmt"
	"testing"
)

func funcA() {
	fmt.Println("A")
}
func funcB() {
	//假如打开了一个数据库连接
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println(err)
			fmt.Println("释放数据库连接")
		}
	}()
	panic("出现严重的错误！！！") //程序崩溃退出
	fmt.Println("B")
}
func funcC() {
	fmt.Println("C")
}

func TestRecover(t *testing.T) {
	funcA()
	funcB()
	funcC()
}
