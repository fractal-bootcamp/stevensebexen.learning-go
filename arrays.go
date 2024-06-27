package main

import "fmt"

func Arrays() {
	var myArr = [5]int32{0, 1, 2, 3, 4}
	var slice = myArr[1 : len(myArr)-1]
	slice = append(slice, 9)  // Overwrite element at myArr[4]
	slice = append(slice, 12) // Do nothing to myArr; slice cap exceeds arr length
	fmt.Printf("%v\n", slice)
	fmt.Printf("%v\n", myArr)
}
