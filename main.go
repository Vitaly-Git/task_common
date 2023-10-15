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

	task_3_deserialize()
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

func task_3_deserialize() {

	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)

}

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Data   ResponseData   `json:"data,omitempty"`
	}

	ResponseHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	ResponseData []ResponseDataItem

	ResponseDataItem struct {
		Type       string                `json:"type"`
		Id         int                   `json:"id"`
		Attributes ResponseDataItemAttrs `json:"attributes"`
	}

	ResponseDataItemAttrs struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}
