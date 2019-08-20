package main

import (
	"fmt"
	"time"
)

// 协程1
func firstGoroutine() {
	fmt.Println("This is first goroutine")
}

// 协程2
func secondGoroutine() {
	time.Sleep(2 * time.Second)
	fmt.Println("This is second goroutine")
}

func main() {
	// 启动协程1,2
	go firstGoroutine()
	go secondGoroutine()

	time.Sleep(1 * time.Second)
	fmt.Println("This is main goroutine")
}
