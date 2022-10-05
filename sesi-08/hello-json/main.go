package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `
	{
		"full_name": "Airell Jordan",
		"email": "airell@mail.com",
		"age": 23
	}
	`

	var result Employee

	var err = json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("ful_name:", result.FullName)
	fmt.Println("email:", result.Email)
	fmt.Println("age:", result.Age)

	var result2 map[string]interface{}

	var err2 = json.Unmarshal([]byte(jsonString), &result2)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}

	fmt.Println("ful_name:", result2["full_name"])
	fmt.Println("email:", result2["email"])
	fmt.Println("age:", result2["age"])

	var temp interface{}

	var err3 = json.Unmarshal([]byte(jsonString), &temp)
	if err3 != nil {
		fmt.Println(err3.Error())
		return
	}

	var result3 = temp.(map[string]interface{})

	fmt.Println("ful_name:", result3["full_name"])
	fmt.Println("email:", result3["email"])
	fmt.Println("age:", result3["age"])

	var jsonString2 = `
	{
		"full_name": "Airell Jordan",
		"email": "airell@mail.com",
		"age": 23
	},
	{
		"full_name": "Ananda RHP",
		"email": "ANANDA@mail.com",
		"age": 23
	}
	`

	var result4 []Employee

	var err4 = json.Unmarshal([]byte(jsonString2), &result4)
	if err4 != nil {
		fmt.Println(err4.Error())
		return
	}

	for i, v := range result4 {
		fmt.Printf("Index %d: %+v\n", i+1, v)
	}
}
