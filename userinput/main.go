package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	//comma ok // err ok // try catch all errors
	//input try and _ => catch
	input, _ := reader.ReadString('\n')

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("input: ", numrating+1)

}
