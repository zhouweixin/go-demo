package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex
var wg sync.WaitGroup

func main() {
	m = new(sync.RWMutex)

	// 读
	go write(1)
	time.Sleep(1 * time.Second)

	// 写
	for i:=3; i<10; i++ {
		go read(i)
	}

	// 读
	go write(2)

	wg.Wait()
	fmt.Println("结束")
}

func read(i int) {
	wg.Add(1)
	defer wg.Done()

	m.RLock()

	println(i, "read start")
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read end")

	m.RUnlock()
}

func write(i int) {
	wg.Add(1)
	defer wg.Done()
	m.Lock()

	println(i, "write start")
	println(i, "writing")
	time.Sleep(5 * time.Second)
	println(i, "write end")

	m.Unlock()
}
