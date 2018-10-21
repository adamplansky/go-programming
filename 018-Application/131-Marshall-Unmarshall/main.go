package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	p := []person{
		person{
			Name:    "Adam Plánský",
			Age:     10,
			Hobbies: []string{"Plavání", "Běhání", "Karate"},
		},
		person{
			Name:    "Eva Nováková",
			Age:     30,
			Hobbies: []string{"Vaření", "Běhání", "Kolo"},
		},
	}
	fmt.Printf("%+v\n", p)
	fmt.Printf(`Marshalling objects: `)
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(`JSON string: `, string(b))

	var unmarshalledPerson []person
	err = json.Unmarshal(b, &unmarshalledPerson)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("UNMARSHALLED go struct: %+v", unmarshalledPerson)
}
