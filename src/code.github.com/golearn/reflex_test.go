package main

import (
	"fmt"
	"reflect"
	"testing"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v; name: %v; kind: %v\n", v, v.Name(), v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func refectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func TestSetValue(t *testing.T) {
	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}

func TestValueOf(t *testing.T) {
	var a float32 = 3.14
	var b int64 = 100
	reflectValue(a) //type is float32, value is 3.140000
	reflectValue(b) //type is int64, value is 100
	//将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c：%T \n", c) //type c: reflect.Value
}

func (s student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t.NumMethod())

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name : %s \n", t.Method(i).Name)
		fmt.Printf("method : %s \n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func TestStruct(t *testing.T) {
	student1 := student{
		Name:  "张三",
		Score: 90,
	}

	tt := reflect.TypeOf(student1)
	fmt.Printf(tt.Name(), tt.Kind()) // student struct
	fmt.Println()
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if field, ok := tt.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	printMethod(student1)
}

type myInt int64

func TestReflex(t *testing.T) {
	var a float32 = 3.14
	reflectType(a) //type:float32
	var b int64 = 100
	reflectType(b) //type:int64

	var c *float32 //指针
	var d myInt    //自定义数据类型
	var e rune     //类型别名
	reflectType(c) //type: kind:ptr
	reflectType(d) //type:myInt kind int64
	reflectType(e) //type int 32 kind: 32

	var p = person1{
		name: "zhangsan",
		age:  19,
	}
	var book1 = book{
		title: "golang",
	}
	reflectType(p)     //type:person1 kind:struct
	reflectType(book1) //type:book kind:struct
}

type person1 struct {
	name string
	age  int
}

type book struct {
	title string
}
