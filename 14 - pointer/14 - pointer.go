package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 42
	//Declaring pointer
	var b *int = &a
	fmt.Println(a, *b)

	a = 27
	//*b: Dereferencing
	fmt.Println(a, *b)

	*b = 14
	fmt.Println(a, *b)

	//Pointer arithmetic
	c := [3]int{1, 2, 3}
	d := &c[0] //Type of d is *int
	e := &c[1]

	//Size of byte
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("c[0]: %T, %d\n", c[0], unsafe.Sizeof(c[0]))
	fmt.Printf("c[1]: %T, %d\n", c[1], unsafe.Sizeof(c[1]))
	fmt.Printf("c[2]: %T, %d\n", c[2], unsafe.Sizeof(c[2]))

	fmt.Printf("%v %p %p\n", c, d, e)

	//Unlike C or C++
	// 	e := &c[1] - sizeof(c[1]) doesn't work

	var ms *myStruct
	ms = &myStruct{foo: 43}
	fmt.Println(ms)
	//&{43} meaning ms is holding the address of an object that
	//has afield with a value of 43

	var ms2 *myStruct
	//Every variable that you declare in Go
	//has an initialization value.
	//ms2, here, is holding nil
	fmt.Println(ms2)
	ms2 = new(myStruct)
	//We need parantheses since deferencing operator actually
	//has a lower precedence than a dot operator
	(*ms2).foo = 57

	//Above line is ugly
	//Compiler understands this
	//Syntatic sugar
	ms.foo = 60

	fmt.Println(ms2)
	fmt.Println((*ms2).foo)

	arr := [3]int{1, 2, 3}
	arr2 := arr
	fmt.Println(arr, arr2)
	arr[1] = 30

	//arr is changed but arr2 is not. arr2 is a copy of arr
	fmt.Println(arr, arr2)

	arr3 := []int{1, 2, 3}
	arr4 := arr3
	fmt.Println(arr3, arr4)
	arr3[1] = 77

	//arr3 is changed, also arr4 is not.
	//arr4 is a copy of arr3
	//Slice is actually a projection of an underlying array
	//Slice doesn't contain the data itself,
	//The slice contains a pointer to the first element
	//that the slice ins pointing to on the underlying array.

	//So, when we work with slices,
	//the internal representation of a slice actually has a pointer.
	//So, when copying we are copying a pointer, not the value

	//SAME THING FOR A MAP
	fmt.Println(arr3, arr4)

}

type myStruct struct {
	foo int
}
