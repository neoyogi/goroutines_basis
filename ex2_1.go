package main

import (
	"fmt"
	"time"
)

func main(){
	defer fmt.Println("end of main goroutine")
	messageCh := make(chan int32)
	go ping(messageCh)
	messageCh <- 1
	go pong(messageCh)
	time.Sleep(100 * time.Millisecond)
}

func ping(messageCh chan int32){
	defer fmt.Println("end of ping goroutine")
	message := <- messageCh
	fmt.Println("output from ping: ", message)
	messageCh <- message + 1
}

func pong(messageCh chan int32){
	defer fmt.Println("end of pong goroutine")
	message :=  <- messageCh
	fmt.Println("output from pong: ", message)
	messageCh <- message + 1
}
