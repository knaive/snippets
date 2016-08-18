package main

import "fmt"

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 10000
	rightMost := make(chan int)
	right := rightMost
	var left chan int
	for i := 0; i < n; i++ {
		left = make(chan int)
		go f(left, right)
		right = left
	}
	rightMost <- 1
	fmt.Println(<-left)
}
