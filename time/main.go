package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hey")
	presentTime := time.Now()
	// fmt.Println(presentTime
	fmt.Println(presentTime.Format("2006/01/02 Sunday"))

	createdDate := time.Date(2002, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

}
