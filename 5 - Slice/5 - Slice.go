package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	fmt.Println(len(a))
	fmt.Println(cap(a))

	//Slices are naturally refrence type
	b := a
	b[1] = 9
	fmt.Println(a)
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d := c[:]
	e := c[3:]
	f := c[:6]
	g := c[3:6]

	c[5] = 777
	g[2] = 999

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	//make function
	//Slice works with the capacity like a vector in C++
	//When the size has to be increased, GO copies all of the existing elements
	//to a new slice with larger size.
	//When the size is small, this copy operation is cheap
	//However, when the size is very big, it can be very expensive
	//That's why we have the third parameter
	makeSlice := make([]int, 3, 100) //slice of integers with 3 lengths with 100 cap
	makeSlice = append(makeSlice, 1, 2, 3, 4, 5, 6)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))

	makeSlice2 := make([]int, 0, 100)
	makeSlice2 = append(makeSlice2, 7, 8, 9, 10)
	fmt.Println(len(makeSlice2))
	fmt.Println(cap(makeSlice2))

	//spread operatir
	makeSlice = append(makeSlice, makeSlice2...)
	makeSlice = append(makeSlice, 0)

	makeSlice = makeSlice[1:]                //removing the first element (index 0)
	makeSlice = makeSlice[:len(makeSlice)-1] //removing the last element (index 0)

	makeSlice = append(makeSlice[:2], makeSlice[5:]...)

	ex1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ex1)
	ex2 := append(ex1[:2], ex1[3:]...)
	fmt.Println(ex1) //1,2,3,4,5,6,6  Adding 4,5,6 at the index of 3
	fmt.Println(ex2)

	//How to remove an element at specific index
	//How to remove an element at specific index
	ex3 := []int{1, 2, 3, 4, 5, 6}
	ex4 := append(append([]int{}, ex3[:2]...), ex3[3:]...)
	fmt.Println(ex3)
	fmt.Println(ex4)
}
