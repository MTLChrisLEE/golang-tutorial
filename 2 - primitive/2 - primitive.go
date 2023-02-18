package main

import "fmt"




func main() {
	//Boolean
	//0 is false
	var isTrue bool = true
	fmt.Printf("%v, %T\n", isTrue, isTrue)
	
	//int
	//int8 -128 to 127
	//int16 -32678 to 32767
	//int32 -2147483648 to 2147483647
	//int64 -9 quintillion to 9 quintillion
	
	//uint8 0-155
	//uint16 0-65535
	//uint32 0-4294967295


	//string in GO is any UTF-8 character, 
	//very powerful but cannot encode every type of character.
	
	s:="This is string"
	//105, uint8 -> In ASCII table, i is 105
	//Stringd in GO are actually aliases for bytes
	fmt.Printf("%v, %T\n", s[2], s[2])	
	fmt.Printf("%v, %T\n", string(s[2]), s[2])

	//Strings are immutable, cannot change a charavter
	
	s2 := []byte(s)	
	fmt.Printf("%v, %T\n", s2, s2)	
	fmt.Println(s2)		
	

	//But why? A of the functions that we are gping to use in GO actually
	//work with slice of bytes and this characteristics give much more flexibility
	
	//For example, if you want to send as a response to a web service call,
	//If you want to send a string back, you can easily convert it to a collection of bytes
	//But if you want to send a file back, 
	//a file on your hard disk is just a collection of bytes too



	//rune
	//rune is just an int32
	//Use single quote
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r) //97, int32 
	


	
}
