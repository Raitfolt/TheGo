package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func remove2(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove(slice []int, i int) []int {
	copy(slice[1:i+1], slice[:i])
	return slice[1:]
}

func reverseOld(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse(slice []int) []int {
	return slice
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d len=%d  cap=%d  \t%v\n", i, len(y), cap(y), y)
		x = y
	}

	s := []int{5, 6, 7, 8, 9}
	fmt.Println(s)
	fmt.Println(remove(s, 2))
	fmt.Println(s)
	fmt.Println(remove2(s, 2))
	fmt.Println(s)
	//123
}
