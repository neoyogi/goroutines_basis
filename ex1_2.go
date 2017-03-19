package main

import (
	"fmt"
	"time"
	"sync"
)

func main()  {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go f1(30, wg)
	wg.Add(1)
	go f2(40, wg)
	fmt.Println("main timeout!")
	wg.Wait()
}

func f1(x int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(30 * time.Millisecond)
	fmt.Println("time out: ", x)
}

func f2(x int, wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(40 * time.Millisecond)
	fmt.Println("time out: ", x)
}
