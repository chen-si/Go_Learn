package main

import (
	"encoding/json"
	"fmt"
)

type Monsters struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

func unserializeStruct() {
	str := "{\"Name\":\"牛魔王\",\"Age\":20,\"Birthday\":\"2011.11.11\",\"Salary\":8000,\"Skill\":\"牛魔角\"}"
	var monster Monsters

	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(str),monster) error:", err)
		return
	}
	fmt.Printf("%T %v\n", monster, monster)

}
func unserializeMap() {
	str := `{"address":"洪崖洞","age":5,"name":"红孩儿"}`

	m := make(map[string]interface{})

	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(str),monster) error:", err)
		return
	}
	fmt.Printf("%T %v\n", m, m)

}

func unserializeSlice() {
	str := `[{"address":"beijing","age":20,"name":"jack"},{"address":"nankang","age":20,"name":"liu"}]`

	s := make([]map[string]interface{}, 5)

	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(str),monster) error:", err)
		return
	}
	fmt.Printf("%T %v\n", s, s)

}

func mains() {
	unserializeStruct()
	unserializeMap()
	unserializeSlice()
}
