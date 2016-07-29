package hello

import "fmt"

func main() {
	// controlStructTest()
	// xs := []float64{1, 2, 3, 4, 5, 6}
	ints := []int{1, 2, 3, 4, 5, 6}
	slice := ints[0:3]
	fmt.Println(sum(slice))
	// fmt.Println(average(xs))

	fmt.Println(half(2))
	fmt.Println(half(4))

	max := findMax(1, 2, 3, 4, 5)
	fmt.Println("Max number is: ", max)

	nextOdd := makeOddGenerator()

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println("fib(10): ", fib(10))

	fib := fibGenerator()
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())

	structTest()
}

func inOutTest() {
	var name string
	fmt.Scanf("%s", &name)
	fmt.Println("Hello go! My name is", name)
}

func controlStructTest() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Round ", i)
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
	arr := [10]int{1, 2, 3, 4, 5, 6}
	for _, value := range arr {
		fmt.Println(value)
	}
}

func average(xs []float64) float64 {
	var total float64 = 0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func sum(xs []int) int {
	total := 0
	for _, value := range xs {
		total += value
	}
	return total
}

func half(x int) (int, bool) {
	x /= 2
	return x, x%2 == 0
}

func findMax(x ...int) int {
	max := 0
	for _, value := range x {
		if value > max {
			max = value
		}
	}
	return max
}

func makeOddGenerator() func() uint {
	nextOdd := uint(1)

	return func() uint {
		defer func() { nextOdd += 2 }()
		return nextOdd
	}
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func fibGenerator() func() int {
	fib1 := 1
	fib2 := 1

	return func() int {
		fib := fib1 + fib2
		defer func() {
			fib1 = fib2
			fib2 = fib
		}()
		return fib
	}
}

type Circle struct {
	x int
	y int
	z int
}

func structTest() {
	var a Circle
	a.x = 1
	a.y = 2
	a.z = 3
	fmt.Println(a)
}
