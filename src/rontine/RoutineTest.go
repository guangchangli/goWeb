package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mu sync.Mutex

func main() {
	//testShare()
	//testChannel()

	arr := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	fmt.Println(arr[3:])
	fmt.Println(arr[:3])
	go sum(arr[3:], c)
	go sum(arr[:3], c)
	x := <-c
	y := <-c
	fmt.Println(x, y, x+y)
}

func testShare() {
	for i := 0; i < 2; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			for j := 0; j < 100; j++ {
				count++
			}
		}()
	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println(count)
}
func testChannel() {
	c := make(chan string)
	c <- "channel"
	//v:=<-c
	fmt.Println(<-c)
	//fatal error: all goroutines are asleep - deadlock!
}

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

/**
channel
*/
