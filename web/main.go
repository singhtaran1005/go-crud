package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO WEB REQUEST")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	} else {
		fmt.Println(response)
	}
	defer response.Body.Close() //callers responsibility to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	} else {
		content := string(databytes)
		fmt.Println(content)
	}
}
