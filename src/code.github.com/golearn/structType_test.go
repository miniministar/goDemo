package main

import (
	"fmt"
	"testing"
)

type person struct {
	name, gender string
	age          int
}

func TestType(t *testing.T) {
	var p person
	p.name = "张三"
	p.gender = "男"
	p.age = 20
	fmt.Println(p.age)

}
