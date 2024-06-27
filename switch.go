package main

import "fmt"

func PrintCascade(a int32) {
	switch a {
	case 0:
		fmt.Println("0")
		fallthrough
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Other")
	}
}

func Switch() {
	const x, y, z int32 = 2, 0, 3
	var vals = [...]int32{x, y, z}
	for _, val := range vals {
		fmt.Printf("Testing %v:\n", val)
		PrintCascade(val)
	}
}
