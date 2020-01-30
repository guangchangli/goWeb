package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	fmt.Println("hello go routine going to sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("go routine awake going to write c")
	c <- true
}

func main() {
	c := make(chan bool)
	fmt.Println("main routine going to call hello routine")
	go hello(c)
	//channel 写入会阻塞
	<-c
	fmt.Println("main receive data out ")
}
