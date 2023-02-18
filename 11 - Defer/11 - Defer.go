package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("start1")
	// fmt.Println("middle1")
	// fmt.Println("end1")

	fmt.Println("============")

	// fmt.Println("start2")
	// defer fmt.Println("middle2")
	// fmt.Println("end2")

	//Defer keyword
	//It actually executes any functions that are passed into it
	//after the function finishes its final statement
	//but before actually returns
	//Here, when go recognizes that the function main exits,
	//it ooks to see if there are nay deferred functons to call.
	//Last-in-First-out, like a stack

	//Example
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal((err))
	}

	//ReadAll function
	//It will take in a stream and that'll parse that out to a string of bytes for you
	robots, err := ioutil.ReadAll(res.Body)

	//In many applications, you might have quite a bit of logic that's invovled
	//that needs that body to be open and continue to work with it.
	//You might make the GET request and open up this resource
	//and it might be dozens of lines later that you actually get around to closing in
	defer res.Body.Close()

	//If you are making a lot of requests and opening a lot of resources witin a loop
	//Then, using the defer keyword to close it might not be the best option
	//since it has to wait for the function to exit

	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("%s", robots)

	a := "start"
	defer fmt.Println(a) //Prints start
	a = "end"
}
