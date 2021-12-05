package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var x = 0
var syncWg sync.WaitGroup
var syncLock sync.Mutex
var syncRWLock sync.RWMutex

func add() {
	for i := 0; i < 50000; i++ {
		syncLock.Lock()
		x += 1
		syncLock.Unlock()
		//time.Sleep(time.Millisecond * 50)
	}
	syncWg.Done()
}
func TestSync(t *testing.T) {
	syncWg.Add(2)
	go add()
	go add()
	syncWg.Wait()
	fmt.Printf("x:%v\n", x)
}

func read() {
	syncRWLock.RLock()
	fmt.Printf("x:%v\n", x)
	time.Sleep(time.Millisecond)
	syncRWLock.RUnlock()
	syncWg.Done()
}

func write() {
	syncRWLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	syncRWLock.Unlock()
	syncWg.Done()
}

func TestWriteRead(t *testing.T) {
	start := time.Now()
	for i := 0; i < 100; i++ {
		syncWg.Add(1)
		go write()
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		syncWg.Add(1)
		go read()
	}
	syncWg.Wait()
	fmt.Printf("%v\n", time.Now().Sub(start))
}
