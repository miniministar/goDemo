package convert

import (
	"fmt"
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	//把数组转换为字符串类型
	i := int32(97)
	str1 := fmt.Sprintf("%d", i) //转换为字符串
	println(str1)

	//字符串转化为整形
	ii, _ := strconv.Atoi(str1)
	fmt.Printf("%#v %T\n", ii, ii)

	str2 := "9000"
	i2, err := strconv.ParseInt(str2, 10, 32)
	if err != nil {
		fmt.Printf("parse int failed ,err:%v \n", err)
	}
	fmt.Printf("%#v %T\n", i2, int(i2))

	//字符串解析为bool值
	boolStr := "true"
	b, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%v %T\n", b, b)

	//字符串解析为float
	floatStr := "3.14159"
	f, _ := strconv.ParseFloat(floatStr, 32)
	fmt.Printf("%v %T\n", f, f)
}
