package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var wg1 sync.WaitGroup

func a() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%v\n", i)
	}
}

func b() {
	defer wg1.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%v\n", i)
	}
}

func TestMaxProcs(t *testing.T) {
	runtime.GOMAXPROCS(2)
	wg1.Add(2)
	go a()
	go b()
	wg1.Wait()
}
