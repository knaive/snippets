package main

import "fmt"

func generator(num ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, value := range num {
			c <- value
		}
		close(c)
	}()
	return c
}

func main() {
	c := generator(1, 2, 3)
	for v, ok := <-c; ok; v, ok = <-c {
		fmt.Println(v)
	}
}
