package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

//goroutine
func hello(i int) {
	fmt.Printf("hello goroutine, %d\n", i)
}

func f() {
	rand.Seed(time.Now().Unix())
	for k := 0; k < 5; k++ {
		i := rand.Int()
		i2 := rand.Intn(10) //0<=x<10
		fmt.Printf("i1:%v, i2:%v\n", i, i2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	println(i)
}

var wg sync.WaitGroup

func TestWaitGoroutineEnd(t *testing.T) {
	//f()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go f1(i)
	}

	//通过WaitGroup计算器等待goroutine都结束
	wg.Wait() //等待wg的计数器减为0
	fmt.Println("main end")
}

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

func TestBasic(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go hello(i) //开启一个单独的goroutine去执行hello函数（任务）
	}
	fmt.Println("main")
	//main函数结束后，由main函数启动的goroutine也都结束了
	//可能会打印 hello,也可能不打印，看hello goroutine执行的速度
	time.Sleep(time.Second)
}
