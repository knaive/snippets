package goroutineTest

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 100; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func g(n int) {
	for i := 0; i < 100; i++ {
		fmt.Println(n, "**", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	go f(0)
	go g(0)
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}
