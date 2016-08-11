package main

import (
    "fmt"
    "container/list"
)

func reverse(l *list.List) *list.List {
    for e:=l.Front(); e!=nil && e.Next()!=nil; {
        next := l.Remove(e.Next())
        l.PushFront(next)
    }
    return l
}

func traverse(l *list.List) {
    for e := l.Front(); e!=nil; e=e.Next() {
        fmt.Println(e.Value)
    }
}

func main() {
    l := list.New()
    for i := 0; i < 6; i++ {
        l.PushBack(i)
    }
    l.PushBack("end")

    traverse(l)

    fmt.Println()
    traverse(reverse(l))
}
