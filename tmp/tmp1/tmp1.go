package main

import "fmt"

//gcd is greatest common divisor
//наибольший общий делитель)
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 1; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	n1 := 65
	n2 := 1375
	fmt.Printf("GCD %d & %d: %d\n", n1, n2, gcd(n1, n2))
	fmt.Println()
	n := 3
	fmt.Printf("%d Fibonacchi's number: %d\n", n, fib(n))
}
