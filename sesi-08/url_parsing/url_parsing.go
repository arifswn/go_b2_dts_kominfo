package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://developer.com:80/hello?name=airell&age=23#history"

	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url string : %s\n", urlString)
	fmt.Printf("url string : %s\n", u.String())
	fmt.Println("")
	fmt.Printf("protocol : %s\n", u.Scheme)
	fmt.Printf("host : %s\n", u.Host)
	fmt.Printf("path : %s\n", u.Path)
	fmt.Println("")
	fmt.Printf("raw query : %s\n", u.RawQuery)
	fmt.Printf("fragment : %s\n", u.Fragment)
	fmt.Println("")
	fmt.Printf("hostname : %s\n", u.Hostname())
	fmt.Printf("port : %s\n", u.Port())
	fmt.Println("")
	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name : %s, age : %s\n", name, age)
}
