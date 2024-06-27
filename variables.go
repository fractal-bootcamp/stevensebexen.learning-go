package main

import "fmt"

func AddTwo(a int32, b int32) int32 {
	return a + b
}

func Variables() {
	var (
		x, y int32 = 890, 1024
		z    int32 = AddTwo(x, y)
	)

	const x0, y0 int32 = 890, 1024
	var z0 int32 = AddTwo(x0, y0)

	const x1, y1 int32 = 0x0187, 0x0a29
	var z1 int32 = AddTwo(x1, y1)

	fmt.Println(z, z0)
	fmt.Printf("%#04x\n", z1)
}
