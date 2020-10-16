package main

import (
	"fmt"
)

func printImage(number int) {
	if number%2 == 0 {
		fmt.Println("The numbers must be odd")
	} else {
		equals := (number / 2) + 1
		add := number + 1
		for i := 1; i < add; i++ {

			for j := 1; j < add; j++ {
				if j == 1 || j == number || i == equals {
					fmt.Print("*  ")
				} else {
					fmt.Print("=  ")
				}
			}
			fmt.Println("\n ")
		}
	}
}

func main() {
	printImage(5)
}
