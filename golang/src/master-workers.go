package main

import (
    "fmt"
    "bufio"
    "os"
)

func worker(substr string, c chan map[string]int) {
    fmt.Println(substr)
    m := make(map[string]int)
    for _, v := range substr {
        s := string(v)
        if _, ok := m[s]; ok {
            m[s] += 1
        } else {
            m[s] = 1
        }
    }
    
    c <- m
}

func master(str string) {
    m := make(map[string]int)
    c := make(chan map[string]int)
    go worker(str[0:len(str)/2], c)
    go worker(str[len(str)/2:], c)

    uniqueCount := 0
    for i:=0; i<2; i++ {
        ret := <-c
        for key, count := range ret {
            if _, ok := m[key]; ok {
                m[key] += count
            } else {
                m[key] = count
            }
            uniqueCount += 1
        }
    }

    for k, v := range m {
        fmt.Println(k, v)
    }

    fmt.Println("count: ", uniqueCount)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)
    text = text[:len(text)-1]
    fmt.Println(text)
    master(text)
}
