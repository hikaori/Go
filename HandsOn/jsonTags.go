package main

import (
	"encoding/json"
	"fmt"
)

type PublicUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func demoJSON() {
	u := PublicUser{ID: 1, Name: "Alicia", Age: 31}
	b, _ := json.Marshal(u)
	fmt.Println(string(b)) // {"id":1,"name":"Alicia","age":31}
}
