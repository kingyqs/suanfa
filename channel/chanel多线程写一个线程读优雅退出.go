package main

import (
	"fmt"
	"sync"
	"time"
)

var msgs = make(chan string)
var doneArr = make(chan struct{}, 3)
var wg = sync.WaitGroup{}

func main() {

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go send(msgs, i, doneArr)
	}

	go receive(msgs)

	wg.Wait()
	fmt.Println("close...")
	close(msgs)
}

func send(ch chan string, k int, done chan struct{}) {
	defer wg.Done()
	fmt.Println("send begin.. k=", k)
	t := time.Duration(k+1) * time.Second
	for {
		<-time.After(t)
		select {
		case <-done:
			fmt.Println("done: k=", k)
			return
		case ch <- fmt.Sprintf("k=%d, msg=%d", k, time.Now().Unix()):
		}
	}
}

func receive(ch chan string) {
	fmt.Println("read begin")
	for i := 0; i < 3; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("read end")
	// 关闭前通知结束
	for i := 0; i < 3; i++ {
		doneArr <- struct{}{}
	}
}
