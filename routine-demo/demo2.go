package main

import "fmt"

func main() {
	done := make(chan int)
	count := 5

	for i := 0; i < count; i++ {
		go func(i int) {
			defer func() {
				done <- 1
			}()
			fmt.Println(i)
		}(i)
	}

	// 协调多个goroutine的时候, channel读的次数要与写的次数相同
	for j := 0; j < count; j++ {
		<- done
	}
	fmt.Println("Over")
}
