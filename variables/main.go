package main

import "fmt"

func main() {

	var username string = "taran"
	fmt.Println(username)
	fmt.Printf("%T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("%T \n", isLoggedIn)

	var smallVal int = 2525
	fmt.Println(smallVal)
	fmt.Printf("%T \n", smallVal)

	//no var keyword

	num := "taran"
	fmt.Println(num)

	//declaring a variable with first letter captial makes it public
}
