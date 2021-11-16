package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs"

func main() {

	fmt.Println("Handling urls")
	fmt.Println(myurl)

	response, _ := url.Parse(myurl)
	fmt.Println(response.RawQuery)

	qparams := response.Query()
	fmt.Printf("%T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println(val)
	}

	//creating new urls
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/hey",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
