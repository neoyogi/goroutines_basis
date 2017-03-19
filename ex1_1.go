package main

import (
	"fmt"
	"time"
)

func main(){
	go f1()
	//f1()
	time.Sleep(1 * time.Second)
	fmt.Println("end of main")
}

func f1() {
	fmt.Println("end of f1")
}
