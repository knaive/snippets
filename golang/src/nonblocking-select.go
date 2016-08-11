package main

import (
    "fmt"
    "time"
)

func main() {
    messages := make(chan string)
    go SendMessage(messages)
    for {
        select {
        case msg := <-messages:
            fmt.Println("received message: ", msg)
        default:
            time.Sleep(time.Millisecond*300)
            fmt.Println("no message received")
        }
    }
}

func SendMessage(c chan string) {
    for {
        c <- "from SendMessage"
        time.Sleep(time.Second*1)
    }
}
