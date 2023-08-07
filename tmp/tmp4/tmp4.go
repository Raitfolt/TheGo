package main

import "fmt"

func main() {
	var ages map[string]int

	fmt.Printf("ages==nil : %v\n", ages == nil)
	fmt.Printf("len(ages)==0 : %v\n", len(ages) == 0)

	ages = make(map[string]int, 0)

	fmt.Printf("%v\n", ages)
	fmt.Printf("%#v\n", ages)

	if _, ok := ages["Bob"]; !ok {
		ages["Bob"] = 20
	}

	fmt.Printf("%v\n", ages)
	fmt.Printf("%#v\n", ages)

	fmt.Printf("%v\n", []string{"123", "234", "word", "hello, world"})
	fmt.Printf("%q\n", []string{"123", "234", "word", "hello, world"})
}
