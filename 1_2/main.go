package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// in loop
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	// use Join
	fmt.Println(strings.Join(os.Args[1:], " "))

	// without format
	fmt.Println(os.Args[1:])

	//1.1
	fmt.Println("Exercise 1.1")
	fmt.Println(strings.Join(os.Args, " "))

	//1.2
	fmt.Println("Exercise 1.2")
	for i, v := range os.Args {
		fmt.Printf("%d: %s", i, v)
	}

	//1.3
	fmt.Println("Exercise 1.3")
	// TODO:  Поэкспериментируйте с измерением разницы времени выполне­ния
	// потенциально неэффективных версий и версии с применением strings.Join.
	// (В разделе 1.6 демонстрируется часть пакета time , а в разделе 11.4 — как написать
	// тест производительности для ее систематической оценки.)
}
