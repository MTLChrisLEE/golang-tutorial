package main

import ("fmt"; 
// "math"
)



const a int16 = 27

func main() {

	//const MyConst if it's going to be exported
	const myConst int = 42;
	fmt.Printf("%v, %T\n", myConst, myConst)


	//math.Sin(1.57) must be executed at runtime, So, it cannot be a constant
	//const trigo float32 = math.Sin(1.57);

	//Constants can be made up of any of the primitive types.


	const a int = 14
	fmt.Printf("%v, %T\n", a, a) // 14, int



	//Enumerated constant
	//iota is scoped in the block
	const (
	 enumerated1 = iota
	 enumerated2 = iota
	 enumerated3 = iota
	)

	//SAME

	// const (
	// 	enumerated1 = iota
	// 	enumerated2 
	// 	enumerated3 
	// )

	fmt.Printf("%v, %T\n", enumerated1, enumerated1) // 0, int
	fmt.Printf("%v, %T\n", enumerated2, enumerated2) // 1, int
	fmt.Printf("%v, %T\n", enumerated3, enumerated3) // 2, int


}
