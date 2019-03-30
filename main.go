package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	aga, err := ioutil.ReadFile("data/aga.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(aga))
}