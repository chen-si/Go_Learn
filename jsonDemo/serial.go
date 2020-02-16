package main

import (
	"encoding/json"
	"fmt"
)

//JavaScript Object Notation

type Monster struct{
	Name string
	Age int
	Birthday string
	Salary float64
	Skill string
}

func testStruct(){
	monster := Monster{
		Name:     "牛魔王",
		Age:      20,
		Birthday: "2011.11.11",
		Salary:   8000.0,
		Skill:    "牛魔角",
	}

	data,err := json.Marshal(monster)
	if err != nil{
		fmt.Println("json.Marshal(monster) error:",err)
		return
	}
	fmt.Println("struct",string(data))

}

func testMap(){
	a := make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 5
	a["address"] = "洪崖洞"
	data,err := json.Marshal(a)
	if err != nil{
		fmt.Println("json.Marshal(a) error:",err)
		return
	}
	fmt.Println("map",string(data))

}

func testSlice(){
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"]="jack"
	m1["age"]=20
	m1["address"]="beijing"
	slice = append(slice,m1)

	m2 := make(map[string]interface{})
	m2["name"]="liu"
	m2["age"]=20
	m2["address"]="nankang"
	slice = append(slice, m2)
	data,err := json.Marshal(slice)
	if err != nil{
		fmt.Println("json.Marshal(a) error:",err)
		return
	}
	fmt.Println("slice",string(data))

}

func testFloat64(){
	a := 19.55

	data,err := json.Marshal(a)
	if err != nil{
		fmt.Println("json.Marshal(a) error:",err)
		return
	}
	fmt.Println("Float64",string(data))
}

func main() {
	//演示将struct map slice 序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
