package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup  // 开箱即用，所以直接声明就好了，没必要用短变量声明
	count := 5

	for i := 0; i < count; i++ {
		// 增加计数器的值
		wg.Add(1)
		go func(i int) {
			// 减少计数器的值, 等同于wg.Add(-1)
			defer wg.Done()

			// do something
			fmt.Println(i)
		}(i)
	}

	// 阻塞当前goroutine, 直到所有子goroutine执行完成
	wg.Wait()
	fmt.Println("Over")
}