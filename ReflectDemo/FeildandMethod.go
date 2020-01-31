package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	Name string
	Age int
	Sex string
}

func (p Person) PrintInfo(){
	fmt.Println("Method:PrintInfo")
	fmt.Printf("name: %s   age: %d   sex: %s\n",p.Name,p.Age,p.Sex)
}

func (p Person) Hello(teacher string,grade int) (string,error){
	fmt.Println("Method:Hello")
	fmt.Printf("Teacher :%s   Grade: %d",teacher,grade)
	return "hahahahahahahahaha..........",nil
}

func main() {
	p := Person{"liu",20,"man"}

	v := reflect.ValueOf(p)

	for i := 0; i < v.NumField() ;i++{
		fmt.Printf("%T   ",v.Field(i))
		fmt.Println(v.Field(i))
	}

	fmt.Println(v.NumMethod())

	for i := 0;i < v.NumMethod(); i++{
		fmt.Printf("%T   ",v.Method(i))
		fmt.Println("Method",i,v.Method(i))
	}
}
