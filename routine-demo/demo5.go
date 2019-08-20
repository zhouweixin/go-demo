package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		// 这个函数会被执行
		once.Do(func() {
			for i := 0; i < 10; i++ {
				fmt.Printf("\r任务[1-%d]执行中...", i)
				time.Sleep(time.Millisecond * 400)
			}
		})
		fmt.Printf("\n任务[1]执行完毕\n")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 300)
		// 这句Do方法的调用会一直阻塞，知道上面的函数执行完毕
		// 然后Do方法里的函数不会执行
		once.Do(func() {
			fmt.Println("任务[2]执行中...")
		})
		// 上面Do方法阻塞结束后，直接会执行下面的代码
		fmt.Println("任务[2]执行完毕")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 300)
		once.Do(func() {
			fmt.Println("任务[3]执行中...")
		})
		fmt.Println("任务[3]执行完毕")
	}()

	wg.Wait()
	fmt.Println("Over")
}
