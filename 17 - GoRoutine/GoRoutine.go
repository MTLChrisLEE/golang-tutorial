package main

import "fmt"
import "time"
import "sync"

func main() {
	//Most of languages use OS threads
	//It meas that they've got an individual function call stack dedicated
	//to the execution of whatever code is handed to that thread.

	//Creation and destruction threads are very expensive, so devs should be conservative

	//In Go, inside the runtime, we've got a scheduler that's going to map these goroutines
	//onto these OS threads for periods of tie.
	//Then, the scheduler will turns with every CPU thread that's available and assign the different go routines.
	//Goroutine is an abstraction of thread, so cheap to create/destory

	//main function is actaully executed in a goroutine itself.
	//So, here we are telling the main function to spawn another goroutine
	//but the application exits as soon as the main function is done.
	//So as soon as it spawsn a goroutine, it finished
	//	go sayHello()
	//	time.Sleep(100 * time.Millisecond)
	



	//Why does this work? Can print Hello or Goodbye
	//Go has a concept called Closure
	//That anonymous function has access to the varialbes in the outer scope
		//var msg = "Hello"
		//go func() {
		//	fmt.Println(msg)
		//}()
		//msg = "GoodBye"

	//time.Sleep(100 * time.Millisecond)
	

	var msg = "Hello2"
	wg.Add(1) //Starts with 1
	//We are copying the value from msg in this case so we are always printing Hello2
	go func(msg string) {
		fmt.Println(msg)
		wg.Done() //Decrement by 1
	}(msg)
	msg = "GoodBye2"


	//time.Sleep(100 * time.Millisecond)


	//Instead of using time, use waitgroup
	wg.Wait() //Wait till waitgroup reaches 0
}


var wg = sync.WaitGroup{}


func sayHello() {
	fmt.Println("Hello")
}




