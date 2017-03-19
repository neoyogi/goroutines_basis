package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	wg := new(sync.WaitGroup)
	messageCh := make(chan int32)
	stopCh := make(chan string)
	duration := time.NewTimer(time.Second * 1)
	wg.Add(1)
	go ping(messageCh, stopCh, wg)
	messageCh <- 1
	wg.Add(1)
	go pong(messageCh, stopCh,  wg)
	wg.Add(1)
	go setExpire(duration, stopCh, wg)
	wg.Wait()
}

func ping(messageCh chan int32, stopCh chan string, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		message := <- messageCh
		fmt.Println("output from ping: ", message)
		messageCh <- message + 1
		<- stopCh
		break
	}
}

func pong(messageCh chan int32, stopCh chan string, wg *sync.WaitGroup){
	defer wg.Done()
	for{
		message :=  <- messageCh
		fmt.Println("output from pong: ", message)
		messageCh <- message + 1
		<- stopCh
		break
	}
}

func setExpire(timer *time.Timer, stopCh chan string,  wg *sync.WaitGroup){
	defer wg.Done()
	<- timer.C
	fmt.Println("Timer expired!")
	stopCh <- "stop"
}
