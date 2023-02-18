package main

import "fmt"

func main() {
	number := 50
	guess := 50

	if guess < number {
		fmt.Println("Too Low")
	} else if guess > number {
		fmt.Println("Too high")
	} else {
		fmt.Println("Holy shit")
	}

	value := 1

	switch value {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Default")
	}

}
