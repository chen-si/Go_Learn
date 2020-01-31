package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := 19.6

	v :=reflect.ValueOf(&f)

	NewValue := v.Elem()

	fmt.Println(NewValue.CanSet())
	NewValue.SetFloat(16.5)

	fmt.Println(f)

	s := []int{1,2,3,4,5}

	v2 := reflect.ValueOf(s)

	fmt.Println(v2.Len())
	for i := 0 ;i < v2.Len();i++{
		NewValue2 := v2.Index(i)
		if NewValue2.CanSet() {
			NewValue2.SetInt(int64(i+6))
		}
	}

	fmt.Println(s)

}