package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person Person
	School string
}

func main() {

	task_1()

	task_2_struct()
}

func task_1() {
	fmt.Println("Test")
}

func task_2_struct() {

	var vitaly Student

	vitaly.School = "MGRU"
	vitaly.Person.Age = 25
	vitaly.Person.Name = "Vitaly"

	ivan := Student{
		School: "VGTU",
		Person: Person{Age: 15, Name: "ivan"},
	}

	fmt.Println(vitaly, ivan)

	result, _ := task_2_struct_marshal(vitaly.Person)

	fmt.Println(string(result))
	fmt.Printf("%v", string(result))

	// city := "Name"
	// fmt.Printf("%T", city[0])

}

func task_2_struct_marshal(person Person) (result []byte, err error) {

	// var result []byte
	result, error := json.Marshal(person)

	// fmt.Println(error)

	return result, error
}
