package main

import "fmt"

type Dragon struct {
	Wings, Legs uint8
}

func Map() {
	dragons := make(map[string]Dragon)
	dragons["Dragon"] = Dragon{Wings: 2, Legs: 4}
	dragons["Wyvern"] = Dragon{Wings: 2, Legs: 2}
	dragons["Kirin"] = Dragon{Wings: 0, Legs: 4}

	for key, value := range dragons {
		fmt.Printf("%s: %d wings, %d legs.\n", key, value.Wings, value.Legs)
	}
}
