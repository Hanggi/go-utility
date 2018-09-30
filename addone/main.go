package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var count int32

func doAdd() {
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		atomic.AddInt32(&count, 1)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(count, " in ", delta)
}

func main() {
	fmt.Println("!!!")

	start := time.Now()
	go doAdd()
	go doAdd()
	end := time.Now()
	delta := end.Sub(start)

	time.Sleep(time.Second * 5)

	fmt.Println(count, " in ", delta)
}
