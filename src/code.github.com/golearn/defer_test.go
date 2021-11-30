package main

import (
	"fmt"
	"testing"
)

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //1.返回值 = y = x = 5, 2.defer修改x, 3.真正返回
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数中的副本
	}(x)
	return 5 // 1.返回值x = 5, 2.defer 修复副本中的x, 3.真正返回
}

func Test01(t *testing.T) {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}
