package main

import (
	"fmt"
	"sync"
	"time"
)

var locker sync.Mutex
var cond = sync.NewCond(&locker)

func main() {
	locker.Lock()
	for i := 0; i < 5; i++ {
		go func(x int) {
			cond.L.Lock()         // 获取锁
			defer cond.L.Unlock() // 释放锁

			cond.Wait() // 等待通知，阻塞当前 goroutine

			// 通知到来的时候, cond.Wait()就会结束阻塞, do something. 这里仅打印
			fmt.Println("111111")
		}(i)
	}

	time.Sleep(time.Second * 1) // 睡眠 1 秒，等待所有 goroutine 进入 Wait 阻塞状态
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发一个通知给已经获取锁的 goroutine

	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 1 秒后下发下一个通知给已经获取锁的 goroutine

	time.Sleep(time.Second * 1)
	cond.Broadcast() // 1 秒后下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 10) // 睡眠 1 秒，等待所有 goroutine 执行完毕
}
