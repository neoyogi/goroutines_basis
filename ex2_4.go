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
	duration := time.NewTimer(time.Second * 5)
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
	defer fmt.Println("ping completed")
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		select {
		case message := <-messageCh:
			fmt.Println("ping: ", message)
			messageCh <- message + 1
		case <-stopCh:
			return
		}
	}
}

func pong(messageCh chan int32, stopCh chan string, wg *sync.WaitGroup){
	defer fmt.Println("pong completed")
	defer wg.Done()
	for {
		time.Sleep(1 * time.Second)
		select {
		case message := <-messageCh:
			fmt.Println("pong: ", message)
			messageCh <- message + 1
		case <-stopCh:
			return
		}
	}
}

func setExpire(timer *time.Timer, stopCh chan string,  wg *sync.WaitGroup){
	defer fmt.Println("end of setExpire")
	defer wg.Done()
	select {
	case <- timer.C:
			fmt.Println("Timer expired!")
			stopCh <- "stop"
			stopCh <- "stop"
	 }
}
