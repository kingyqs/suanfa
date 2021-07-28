package main

import (
	"fmt"
	"time"
)

var msgs = make(chan int)
var done = make(chan struct{})

func main() {
	go send(msgs)
	go receive(msgs)

	time.Sleep(5 * time.Second)
}

func send(ch chan int) {
	for {
		<-time.After(1 * time.Second)
		select {
		case <-done:
			fmt.Println("done")
			return
		case ch <- time.Now().Second():
		}
	}
}

func receive(ch chan int) {
	<-time.After(1 * time.Second)
	msg := <-ch
	fmt.Println(msg)
	// 关闭前通知结束
	done <- struct{}{}
	close(ch)
}
