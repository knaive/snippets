package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fanin(input1, input2 chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func fanInSelect(c1, c2 chan string) <-chan string {
	cout := make(chan string)

	go func() {
		for {
			select {
			case s := <-c1:
				cout <- s
			case s := <-c2:
				cout <- s
			}
		}
	}()

	return cout
}

func main() {
	input1 := make(chan string)
	input2 := make(chan string)

	go func() {
		for {
			input1 <- "Jack"
			time.Sleep(time.Millisecond * 100 * time.Duration(rand.Intn(10)))
		}
	}()

	go func() {
		for {
			input2 <- "Tom"
			time.Sleep(time.Millisecond * 100 * time.Duration(rand.Intn(10)))
		}
	}()

	// c := fanin(input1, input2)
	c := fanInSelect(input1, input2)

	for {
		fmt.Println(<-c)
	}
}
