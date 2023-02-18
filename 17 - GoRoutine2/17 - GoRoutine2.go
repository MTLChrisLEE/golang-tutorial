package main

//Using this mutex and waitgorup in this example destorys the benefits of concurrency

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	//Result is inconsistent without mutex. Time to use mutex
	for i := 0; i < 10; i++ {
		wg.Add(2) //We are adding two goroutines in each loop
		m.RLock() //Locking outside of goroutine (in a single context)
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()

	//A mutex is basically a lock that the application is going to honor
	//A simple mutex is simply locked or unlocked.
	//If someting tries to manipulate that value, it has to wait until the mutex unlocks

	//With RWMutex, we've basically said
	//"as many things as want to can read this data" but only one ca write it at a time
	//If anything is reading, then we can't write to it at all.

}

func sayHello() {
	// m.RLock()
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}
