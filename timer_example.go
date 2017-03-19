package main

import (
	"time"
	"fmt"
)
func main()  {
	timer1 := time.NewTimer(time.Second * 1)
	go func() {
		<- timer1.C
		fmt.Println("timer expired!")
	}()
	time.Sleep(2 * time.Second)
}
