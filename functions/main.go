package main

import "fmt"

func main() {
	res := add(3, 5)
	fmt.Println("result is = ", res)
}

func add(valone int, valtwo int) int {
	return (valone + valtwo)
}

//function having all data members as integers
// func proadd(values ...int) int {
// }
