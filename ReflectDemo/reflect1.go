// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Person1 struct{
// 	name string
// 	age int
// 	sex string
// }

// func main() {
// 	var f = 19.6 //float64
// 	v :=reflect.ValueOf(f)
// 	t :=reflect.TypeOf(f)

// 	fmt.Printf("%T ,%V\n\n\n",t,t)

// 	fmt.Println("v.Kind()",v.Kind())
// 	fmt.Println("v.Type()",v.Type())

// 	//接口的类型断言
// 	convertValue := v.Interface().(float64)
// 	fmt.Printf("%T %f\n",convertValue,convertValue)

// 	fmt.Println("------------------------------------------------")

// 	p := Person1{"liu",20,"man"}

// 	v2 := reflect.ValueOf(p)
// 	t2 := reflect.TypeOf(p)

// 	fmt.Println(t2)

// 	fmt.Println("v2.Kind()",v2.Kind())
// 	fmt.Println("v2.Type()",v2.Type())

// }
