package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("========")

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(j)
	}

	fmt.Println("========")

	i := 0
	for ; i < 5; i++ {
		fmt.Printf("i: %v\n", i)
	}

	fmt.Println("========")

	j := 0
	for j < 5 {
		fmt.Printf("j: %v\n", j)
		j++
	}

	fmt.Println("========")

Loop:
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			fmt.Printf("i*j %v\n", a*b)
			if a*b >= 3 {
				break Loop
			}

		}
	}

	//Looping over collections have similar syntax
	// arrays, slices, maps, strings, channels
	// for k,v := range collectionName {}

	fmt.Println("========")
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

}
