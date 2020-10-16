package main

import "fmt"

type rata interface {
	rataKanan()
	rataKiri()
}

type star struct {
	value int
}

func main() {
	var number int
	fmt.Print("Input number: ")
	fmt.Scanf("%d", &number)

	var pattern rata
	pattern = star{number}
	pattern.rataKiri()
	fmt.Println("\n")
	pattern.rataKanan()
}

func (c star) rataKiri() {
	for i := 1; i <= c.value; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func (c star) rataKanan() {
	for i := 1; i <= c.value; i++ {
		for j := 1; j <= i; j++ {
			if j > 1 {
				fmt.Print(" ")
			}
		}
		for k := i; k <= c.value; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
