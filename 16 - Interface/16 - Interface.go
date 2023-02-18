package main

import "fmt"

//Single method interface is very common in the language
//Very powerful and define a very specific behavior

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

}

// Interface describes behaviors
// Methode definition
type Writer interface {
	Write([]byte) (int, error)
}

// No need keyword such as "implements"
// No need to be a struct
// Any type that can have a method assicated with it can implement an interface
// Any type can have methods associated with it.
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

// When I'm implemeting an interface, if I use a value type
// the methods that implement the interface have to all have value receivers.

// If I'm implementing the interface with a pointer,
// then I just have to have the methods there, regardless of the receiver type
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

//Best Practices
// - Use many, small interfaces
// 	- Single method interfaces are some of the most powerful and flexible
// - Don't export interfaces for types that will be consumed
// - Do export interfaces for types that will be used by package
