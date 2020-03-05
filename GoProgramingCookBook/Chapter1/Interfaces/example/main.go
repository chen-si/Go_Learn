package main

import (
	"bytes"
	"fmt"

	"GoProgramingCookBook/Chapter1/Interfaces"
)

func main() {
	in := bytes.NewReader([]byte("example"))
	out := &bytes.Buffer{}
	fmt.Print("std out on Copy = ")
	if err := Interfaces.Copy(in, out); err != nil {
		panic(err)
	}
	fmt.Println("out bytes buffer = ", out.String())
	fmt.Print("stdout on Pipeexample")
	if err := Interfaces.PipeExamples(); err != nil {
		panic(err)
	}
}
