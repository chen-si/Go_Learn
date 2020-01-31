package main

import (
	"fmt"
	"reflect"
)

type Person2 struct {
	Name string
	Age  int
	Sex  string
}

func (p Person2) PrintInfo(){
	fmt.Println("Method:PrintInfo")
	fmt.Printf("name: %s   age: %d   sex: %s\n",p.Name,p.Age,p.Sex)
}

func (p Person2) Hello(teacher string,grade int) (string,error){
	fmt.Println("Method:Hello")
	fmt.Printf("Teacher :%s   Grade: %d\n",teacher,grade)
	return "hahahahahahahahaha..........",nil
}

func fun1(){
	fmt.Println("This is fun1......")
}

func main() {
	p := Person2{"liu",20,"man"}

	v := reflect.ValueOf(p)

	MethodValue1 := v.MethodByName("PrintInfo")

	MethodValue1.Call(nil)

	MethodValue2 := v.MethodByName("Hello")
	CallValue2 := []reflect.Value{reflect.ValueOf("zen"),reflect.ValueOf(100)}

 	ResultValue := MethodValue2.Call(CallValue2)

 	fmt.Println(ResultValue)

 	FuncValue := reflect.ValueOf(fun1)

 	FuncValue.Call(nil)
}
