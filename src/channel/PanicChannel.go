package main

func main() {
	//dealLock()

	//create send only channel
	ch := make(chan<- int)
	go singleChannel(ch)

}

func dealLock() {
	c := make(chan int)
	c <- 1
}
func singleChannel(c chan<- int) {
	c <- 10
}
