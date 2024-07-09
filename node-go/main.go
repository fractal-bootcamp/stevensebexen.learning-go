package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cheese struct {
	Name       string `json:"name"`
	Cheesiness int    `json:"cheesiness"`
	Message    string `json:"message"`
}

func main() {
	cheeseList := [...]Cheese{
		{"Mozzarella", 8, "A light, stringy cheese with a wonderful, light flavor."},
		{"Cheddar", 9, "A solid all-arounder cheese with excellent macaroni potential."},
	}

	j, err := json.Marshal(cheeseList)
	if err != nil {
		log.Fatalf("Oh nooo there was a problem: %s\n", err)
	}

	fmt.Printf("%s\n", string(j))
}
