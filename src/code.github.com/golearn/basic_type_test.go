package main

import (
	"fmt"
	"testing"
)

var str string = `
多行字符串
，第二行
`

func TestBasic(t *testing.T) {
	//字符串长度
	fmt.Println(len(str), fmt.Sprint(str, "拼接字符串"), str)
	var inta int = 32
	fmt.Printf("%d", inta)
	fmt.Printf("%T", inta)
	fmt.Println()
	var float32a float32 = 32.15
	fmt.Printf("%f", float32a)
	fmt.Printf("%T", float32a)
	fmt.Println()
	var booleana bool = true
	fmt.Printf("%t", booleana)
	fmt.Printf("%T", booleana)
	fmt.Println()
	var stra string = "abc"
	fmt.Printf("%s", stra)
	fmt.Printf("%T", stra)
	fmt.Println()
}
