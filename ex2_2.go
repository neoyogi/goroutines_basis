package main

import (
	"fmt"
	"sync"
)

func main(){
	wg := new(sync.WaitGroup)
	messageCh := make(chan int32)
	wg.Add(1)
	go ping(messageCh, wg)
	messageCh <- 1
	wg.Add(1)
	go pong(messageCh, wg)
	wg.Wait()
}

func ping(messageCh chan int32, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		message := <- messageCh
		fmt.Println("output from ping: ", message)
		messageCh <- message + 1
	}
}

func pong(messageCh chan int32, wg *sync.WaitGroup){
	defer wg.Done()
	for{
		message :=  <- messageCh
		fmt.Println("output from pong: ", message)
		messageCh <- message + 1
	}
}
