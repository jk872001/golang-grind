package main

import "fmt"

const age = 32

func main(){
	const (
		port = 5000
		url = "localhost"
	)

	fmt.Println(port,url)
}
