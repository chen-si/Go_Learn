package bytestrings

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//SearchString show a number of methods for Search a string
func SearchString() {
	s := "this is a test"
	//return true beacuse s contains the word this
	fmt.Println(strings.Contains(s, "this"))
	//return true because s contains the letter a would also match if it contained b or c
	fmt.Println(strings.ContainsAny(s, "abc"))
	//return true because s start with this
	fmt.Println(strings.HasPrefix(s, "this"))
	//return true because s ends with test
	fmt.Println(strings.HasSuffix(s, "test"))

}

//ModifyString modifies a string in a number of ways
func ModifyString() {
	s := "simple string"
	//prints [simple string]
	fmt.Println(strings.Split(s, " "))
	//prints "Simple String"
	fmt.Println(strings.Title(s))
	//prints "simple string";all tailing and leading white space is removed
	s = " simple string   "
	fmt.Println(strings.TrimSpace(s))
}

//StringReader domonstrates how to create a io.Reader with a string
func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)
	//print s on stdout
	io.Copy(os.Stdout, r)
}
