package main

import "fmt"

// a buggy insert
func Insert(slice []int, index, value int) []int {
    slice = slice[0:len(slice)+1]
    copy(slice[index+1:], slice[index:])
    slice[index] = value
    return slice
}

func Extend(slice []int, element int) []int {
    n := len(slice)
    if n == cap(slice) {
        newSlice := make([]int, n, 2*n+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:n+1]
    slice[n] = element
    return slice
}

func DullAppend(slice []int, items ...int) []int {
    for _, item := range items {
        slice = Extend(slice, item)
    }
    return slice
}

func Append(slice []int, items ...int) []int {
    n := len(slice)
    total := n + len(items)
    if total > cap(slice) {
        capacity := total*3/2 + 1
        newSlice := make([]int, total, capacity)
        copy(newSlice, slice)
        slice = newSlice
    }
    copy(slice[n:], items)
    return slice
}

func main() {
    fmt.Println("Insert: ")
    slice := []int{1,2,3,4,5}
    slice = slice[:4]
    fmt.Println(slice)
    Insert(slice, 2, 6)
    fmt.Println(slice)

    fmt.Println("\nExtend: ")
    slice = []int{1,2,3,4,5}
    Extend(slice, 6)
    fmt.Println(slice, len(slice), cap(slice))

    fmt.Println("\nDullAppend: ")
    slice = []int{1,2,3,4,5}
    fmt.Println(slice)
    slice = DullAppend(slice, 6, 7, 8)
    fmt.Println(slice)

    fmt.Println("\nDullAppend: ")
    slice1 := []int{1,2}
    slice2 := []int{3,4,5}
    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(DullAppend(slice1, slice2...))

    fmt.Println("\nAppend: ")
    slice1 = []int{1,2}
    slice2 = []int{3,4,5}
    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(Append(slice1, slice2...))

    fmt.Println("\nAppend a slice to itself: ")
    slice1 = []int{1,2}
    fmt.Println(slice1)
    fmt.Println(Append(slice1, slice1...))
}
