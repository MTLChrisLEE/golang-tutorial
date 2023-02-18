//ALL VARIABLES MUST BE USED

//VISIBILTY RULES
//  Lower case first letter for package scope
//  Upper case first letter to export
//  No Private Scope

package main

import (
	"fmt"
	"strconv"
)

// On the package level, full declaration syntax is necessary
var wtv int = 42

// block of variables
var (
	actorName    string = "Elisabeth Olsen"
	companion    string = "Wtv"
	doctorNumber int    = 3
	season       int    = 11
)

// Shadowing
var j int = 27

func main() {
	//There are 3 ways to declare a variable
	var i int //When not ready to initialize
	i = 60
	fmt.Println(i)
	i = 120
	fmt.Println(i)

	var j int = 50
	//Printing 50 because of shadowing
	//Shadowing: Taking precedence of inner most variable

	fmt.Println(j)
	//var j int = 51 won't work since it's already declared on line 50

	j = 100
	fmt.Println(j)
	fmt.Printf("%v, %T", j, j)

	k := 40
	fmt.Println(k)
	fmt.Printf("%v, %T\n", k, k)

	fmt.Println(actorName)

	//Type casting
	var integer int = 42
	fmt.Printf("%v, %T\n", integer, integer)
	var floatNumber float32
	//floatNumber = integer NOT WORKING
	floatNumber = float32(integer)
	fmt.Printf("%v, %T\n", floatNumber, floatNumber)

	var randomNumber int = 35
	fmt.Printf("%v, %T\n", randomNumber, randomNumber)

	var castedString string
	castedString = string(randomNumber) //The value of castedString here is #, ASCII
	fmt.Printf("%v, %T\n", castedString, castedString)
	castedString = strconv.Itoa(randomNumber)

	fmt.Printf("%v, %T\n", castedString, castedString)
}
