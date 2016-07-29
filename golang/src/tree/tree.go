package tree

import (
    "fmt"
    "math/rand"
)
type Tree struct {
    Left *Tree
    Value int
    Right *Tree
}

func New(k int) *Tree {
    var t *Tree
    for _, v := range rand.Perm(10) {
        t = insert(t, (1+v)*k)
    }
    return t
}

func insert(t *Tree, k int) *Tree {
    if t == nil {
        return &Tree{nil, k, nil}
    }
    if t.Value > k {
        t.Left = insert(t.Left, k)
    } else {
        t.Right = insert(t.Right, k)
    }
    return t
}

func Traverse(t *Tree) {
    if t == nil {
        return
    }
    fmt.Println(t.Value)
    Traverse(t.Left)
    Traverse(t.Right)
}
