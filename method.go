package main

import "fmt"

type Duck struct {
	name        string
	catchphrase string
	numFish     uint
}

func (d Duck) Quack() {
	fmt.Printf("%s quacks: %q\n", d.name, d.catchphrase)
}

func (d *Duck) CatchFish(numFish uint) {
	d.numFish += numFish
}

func TakeFish(d *Duck, numFish uint) {
	d.numFish -= numFish
}

func Method() {
	var doug Duck = Duck{name: "Doug", catchphrase: "Hey, what's up?", numFish: 0}
	doug.Quack()
	doug.CatchFish(2)
	fmt.Printf("Doug has caught %d fish.\n", doug.numFish)
	TakeFish(&doug, 1)
	fmt.Printf("Doug has paid taxes and now has %d fish.\n", doug.numFish)
}
