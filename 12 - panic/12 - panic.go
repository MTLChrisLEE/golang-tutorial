package main

import "net/http"

func main() {
	//In Go, we don't have exceptions.
	//When you try to open a file that doen's exist
	//Go returns an error instead of throwing an exception
	//Panic happens AFTER the defer happens

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("Something had happened")
	// fmt.Println("end")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(([]byte("Hello Go!")))
	})

	//ListenAndServe function doesn't have an opinion on if that's a panicking situation or not
	//Because it just tried to execute something.
	// It's going to return an error value saying that didn't work

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic((err.Error()))
	}

}
