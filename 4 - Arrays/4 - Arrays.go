package main

import "fmt"

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 100
	fmt.Printf("Grades: %v, %v, %v\n", grade1, grade2, grade3)

	// grades := [3]int
	//grades := [...]int{97,85,100} Create an arrau that's just large enough to old the data
	grades := [3]int{97, 85, 100}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grade: %v\n", grades[0])
	fmt.Printf("Lenth: %v\n", len(grades))

	//Arrays are actually considered values in Go.
	//In a lot of languages, when you create an array,
	//it's actually pointing to the values.
	//So if you pass things around, you're actually passing around the same underlying data.

	//But un GO, that's not true. When you copy an array, you are actually creating a literal copy.
	//So, it's not pointing to the same underying data, it is pointing to a different set of data.
	a := [...]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	c := [...]int{1, 2, 3}
	d := &c //pointer
	c[1] = 5
	fmt.Println(c)
	fmt.Println(d)

}
