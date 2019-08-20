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
		defer func() {
			if p := recover(); p != nil {
				fmt.Printf("PANIC: %v\n", p)
				// 下面的语句会给once变量替换一个新的Once值，这样下面的第二个任务还能被执行
				//once = sync.Once{}
			}
		}()

		once.Do(func() {
			fmt.Println("开始执行参数函数，紧接着会引发panic")
			panic(fmt.Errorf("主动引发了一个panic"))  // panic之后就去调用defer了
			fmt.Println("参数函数执行完毕")  // 这行不会执行，后面的都不会执行
		})
		fmt.Println("Do方法调用完毕")  // 这行也不会执行
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 500)
		once.Do(func() {
			fmt.Println("第二个任务执行中...")
			time.Sleep(time.Millisecond * 800)
			fmt.Println("第二个任务执行结束")
		})
		fmt.Println("第二个任务结束")
	}()

	wg.Wait()
	fmt.Println("Over")
}
