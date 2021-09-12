package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var signChan chan bool = make(chan bool, 1)

func f() {
	defer wg.Done()
LOOP:
	for {
		fmt.Println("xxxxx")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-signChan:
			break LOOP
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 2)
	signChan <- true
	wg.Wait()
	fmt.Println("结束了")
}
