package main

import (
	"GoProgramingCookBook/Chapter1/bytestrings"
)

func main() {
	err := bytestrings.WorkWithBuffer()
	if err != nil {
		panic(err)
	}
	//each of this print to stdout
	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
