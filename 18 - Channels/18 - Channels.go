package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// Channels are designed to synchronize data transmission
// between multiple goroutines.
func main() {

	//Channel is strongly-typed
	// ch := make(chan int)
	// wg.Add(2)

	//Receiving goroutine
	// go func() {
	// 	i := <-ch //Receivin a value from the channel and assigning it to i
	// 	fmt.Println(i)
	// 	ch <- 27
	// 	wg.Done()
	// }()

	// //Sending goroutine
	// go func() {
	// 	//This line is actually going to pause the execution of this goroutine
	// 	//until there's a space available in the channel.
	// 	//By default, we're working with unbuffered channels
	// 	//which means only one message can be in the channel
	// 	ch <- 42
	// 	fmt.Println((<-ch))
	// 	wg.Done()
	// }()

	//////////////////////////////

	// //Receiving goroutine
	// go func(ch <-chan int) { //making receive only
	// 	i := <-ch //Receivin a value from the channel and assigning it to i
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)

	// //Sending goroutine
	// go func(ch chan<- int) {
	// 	//This line is actually going to pause the execution of this goroutine
	// 	//until there's a space available in the channel.
	// 	//By default, we're working with unbuffered channels
	// 	//which means only one message can be in the channel
	// 	ch <- 42
	// 	wg.Done()
	// }(ch)

	//////////////////////////

	// ch2 := make(chan int, 50)

	// go func(ch2 <-chan int) { //making receive only
	// 	i := <-ch2 //Receivin a value from the channel and assigning it to i
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch2)

	// //Sending goroutine
	// go func(ch2 chan<- int) {
	// 	//Can send two messages
	// 	ch2 <- 42
	// 	ch2 <- 27
	// 	wg.Done()
	// }(ch2)

	// wg.Wait()

	//////////////////

	ch3 := make(chan int, 50)
	wg.Add(2)

	go func(ch3 <-chan int) { //making receive only
		for {
			if i, ok := <-ch3; ok { //ok lets us know if a channel is open or not
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch3)

	//Sending goroutine
	go func(ch3 chan<- int) {
		//Can send two messages
		ch3 <- 42
		ch3 <- 27
		//Need to close to let readers detect that.
		//w/o close, readers iterate infinitely
		//Be careful with closing a channel!!! After closing, cannot send a msg
		//Cannot even undo clsing a channel
		close(ch3)
		wg.Done()
	}(ch3)

	wg.Wait()

}
