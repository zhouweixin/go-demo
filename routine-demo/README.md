# goroutine协程

## 1 协程的创建与使用: demo1.go

```go
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
```
当主协程结束时, 子协程自动结束, 即使任务没有完成.

## 2 利用管道channel等待子协程执行: demo2.go

```go
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
```

## 3 利用sync.WaitGroup等待子协程执行: demo3.go

sync.WaitGroup的3个方法:
- Add(int)
- Done()
- Wait()


```go
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
```

## 4 
