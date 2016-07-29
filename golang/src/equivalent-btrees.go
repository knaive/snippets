package main

import (
    "fmt"
    "tree"
)

func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        ch <- 0
        return
    }
    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
    count := 10
    c1 := make(chan int)
    c2 := make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for i:=0; i<count; i++ {
        if <-c1 != <-c2 {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
