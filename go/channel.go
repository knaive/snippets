package channelTest

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func pinger(c chan string) {
	var msg string = "ping"
	for i := 0; ; i++ {
		fmt.Println("pinger: ", msg)
		c <- msg
	}
}

func ponger(c chan string) {
	var msg string = "pong"
	for i := 0; ; i++ {
		fmt.Println("ponger: ", msg)
		c <- msg
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("printer: ", msg)
		time.Sleep(time.Second)
	}
}
