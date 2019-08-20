package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter uint32
	var once sync.Once

	once.Do(func() {
		atomic.AddUint32(&counter, 1)
	})
	fmt.Println("counter:", counter)

	// 这次调用不会被执行
	once.Do(func() {
		atomic.AddUint32(&counter, 2)
	})
	fmt.Println("counter:", counter)
}
