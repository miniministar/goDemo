package test

import (
	"reflect"
	"testing"
)

// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted: %v, got: %v\n", want, got)
	}
}

func Test2Split(t *testing.T) {
	got := Split("a::b::c", "::")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted: %v, got: %v\n", want, got)
	}
}
