package main

import "fmt"

func main() {
	months := [...]string{1: "Январь", "Февраль", "Март", "Апрель", "Май", "Июнь",
		"Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s находится в обоих срезах\n", s)
			}
		}
	}
	fmt.Printf("Q2 len: %d cap %d\n", len(Q2), cap(Q2))
	fmt.Printf("summer len: %d cap %d\n", len(summer), cap(summer))
}
