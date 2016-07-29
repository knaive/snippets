package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            fmt.Println("fibonacci written")
            x, y = y, x+y
        case <- quit:
            fmt.Println("exit...")
            return
        }
    }
    close(c)
}

func main() {
    size := 10
    c := make(chan int, size)
    quit := make(chan int)
    go func() {
        for i:=0; i<10; i++{
            v := <-c
            fmt.Println("anonymous")
            fmt.Println(v)
        }
        quit <- 0
    }()

    fibonacci(c, quit)
}
