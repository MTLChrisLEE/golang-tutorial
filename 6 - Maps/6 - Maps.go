package main

import "fmt"

func main() {

	//JUST LIKE SLICES
	//REFERENCE!

	statePopulations := map[string]int{
		"CA": 40000000,
		"TX": 28000000,
	}

	//Adding a pair
	//The order is NOT guaranteed
	statePopulations["NY"] = 20000000
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["CA"])
	fmt.Println(len(statePopulations))

	fmt.Println("=========================")

	statePopulations2 := make(map[string]int)
	statePopulations2["CA"] = 40000000
	statePopulations2["TX"] = 28000000
	fmt.Println(statePopulations2)
	delete(statePopulations2, "TX")
	fmt.Println(statePopulations2)
	fmt.Println(statePopulations2["TX"]) //Returns 0

	fmt.Println("=========================")
	pop, isInMap := statePopulations2["TX"]
	fmt.Println(pop)
	fmt.Println(isInMap)

	fmt.Println("=========================")
	//Conventional to use 'ok'
	_, ok := statePopulations2["TX"]
	fmt.Println(ok)

}
