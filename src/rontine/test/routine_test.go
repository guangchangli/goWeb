package test

import (
	"fmt"
	"testing"
	"time"
)

func hello() {
	fmt.Println("hello world go routine")
}
func Test_Routine(t *testing.T) {
	go hello()
	time.Sleep(time.Second)
	fmt.Println("main routine out")
}
