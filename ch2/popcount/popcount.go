package main

import (
	"fmt"
	"time"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var result int
	for i := 0; i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

func PopCountShift(x int) int {
	var result int
	for x > 0 {
		result += LeftBitCheck(x)
		x = x >> 1
	}
	return result
}

func PopCountBitReset(x int) int {
	var result int
	for x > 0 {
		RightBitReset(&x, &result)
	}
	return result
}

func RightBitReset(x, res *int) {
	*x = *x & (*x - 1)
	*res++
}

func LeftBitCheck(x int) int {
	return x & 1
}

func main() {

	start := time.Now()
	for i := 0; i < 65432; i++ {
		_ = PopCount(uint64(i))
	}
	fmt.Printf("PopCount: %d microseconds elapsed\n", time.Since(start).Microseconds())

	start = time.Now()
	for i := 0; i < 65432; i++ {
		_ = PopCountLoop(uint64(i))
	}
	fmt.Printf("PopCountLoop: %d microseconds elapsed\n", time.Since(start).Microseconds())

	start = time.Now()
	for i := 0; i < 65432; i++ {
		_ = PopCountShift(i)
	}
	fmt.Printf("PopCountShift: %d microseconds elapsed\n", time.Since(start).Microseconds())

	start = time.Now()
	for i := 0; i < 65432; i++ {
		_ = PopCountBitReset(i)
	}
	fmt.Printf("PopCountBitReset: %d microseconds elapsed\n", time.Since(start).Microseconds())

}
