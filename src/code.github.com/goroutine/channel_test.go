package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var wgChannel sync.WaitGroup

func TestChannelDefine(t *testing.T) {
	var a []int
	fmt.Printf("%v\n", a)
	var b chan int        //需要指定通道中元素的类型
	fmt.Printf("%v\n", b) //nil
	b = make(chan int)    //不带缓冲区通道的初始化
	//b <- 10		//无人接收，hang住了，死锁
	wgChannel.Add(1)
	go func() {
		defer wgChannel.Done()
		x := <-b
		fmt.Printf("后台goroutine从通道b中取到了%v\n", x)
	}()
	b <- 10
	fmt.Println("10发送到通道b中了...")
	//b = make(chan int, 16)	//待缓冲区的通道的初始化
	fmt.Println(b)
	wgChannel.Wait()
}

func TestBuffer(t *testing.T) {
	b := make(chan int, 16) //待缓冲区的通道的初始化
	//wgChannel.Add(1)
	//go func() {
	//	defer  wgChannel.Done()
	//	x:=<- b
	//	fmt.Printf("后台goroutine从通道b中取到了%v\n", x)
	//}()
	b <- 10
	fmt.Println("10发送到通道b中了...")
	b <- 20
	fmt.Println("20发送到通道b中了...")
	x := <-b
	fmt.Println("从通道b中取到", x)
	close(b)

}
