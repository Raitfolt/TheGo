package main

import (
	"fmt"
	"strings"
)

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func squaresF() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {

	defer func() {
		if recover() != nil {
			fmt.Println("Was panic")
		}
	}()

	f := square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	f2 := product
	fmt.Println(f2(3, 5))

	fmt.Println(strings.Map(add1, "Hello, world"))

	fmt.Println(strings.Map(func(r rune) rune { return r + 2 }, "Hello, world"))

	sF := squaresF()
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
	fmt.Println(sF())
}

func add1(r rune) rune {
	return r + 1
}
