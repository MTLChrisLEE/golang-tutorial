package main

import "fmt"

func main() {

	sayMessage("Hello GO")

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(greeting, &name)
	fmt.Println(name)

	sum(1, 2, 3, 4, 5)
	fmt.Println(sum2(1, 2, 3, 4, 5))
	fmt.Println(sum3(1, 2, 3, 4, 5))

	d, err := divide(5.0, 1.3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	//Anonymous function
	func() {
		fmt.Println("HELLO GOGOGOG")
	}() //This () invovkes the anonymous function

	//A great use of anomoymous function is in a for loop

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	anonyFunc := func() {
		fmt.Println("I am anonymouns func")
	}
	anonyFunc()

	var divideAnonyFunc func(float64, float64) (float64, error)
	divideAnonyFunc = func(aFloat, bFloat float64) (float64, error) {
		if bFloat == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return aFloat / bFloat, nil
	}

	resultDivide, err := divideAnonyFunc(3.4, 4.2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resultDivide)

	g := greeter{
		greeting: "Hello World,",
		name:     "I am Go",
	}
	g.greet()

	fmt.Println(g.name)
	g.greetObject()
	fmt.Println(g.name)

}

// Starts with 'func' keyword
// Depending on you have an uppercase or lowercase
// determines the visibility of that function
// lower case: private
func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting string, name *string) {
	fmt.Println(greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// Passing a large data structure in the value of that data structure
// is going to cause that entire data data structure to be copied

// Variadic parameter
// Telling Go runtime to take in all of the last argumetns and
// wrap them up to a slice

// When using a variadic parameter,
// you can only have one (No more than 2 variadic parmeter) and it has to be the last one
// func sum(msg string, values ... int) would work
func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

// return type
func sum2(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return result
}

// returning a pointer type
// This is safe in Go
// Go reconginzes that the function is retuning a local pointer
// It automatically promotes it to heap
func sum3(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}

	return &result
}

// names return values
// NOT recommended to use
func sum4(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// Returning multiple values
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// Methods
// A method is just a function with a receiver argument.
// It is declared with the same syntax with the addition of the receiver.

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "" //Doesn't change the value here. Working on a copy
}

func (g *greeter) greetObject() {
	fmt.Println(g.greeting, g.name)
	(*g).name = "" //Changes the value here. Working on the underlying data
}
