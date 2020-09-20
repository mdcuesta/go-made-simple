package main

import "fmt"

type Boat struct {
	Name string

	occupants []string
}

func (b *Boat) AddOccupant(name string) *Boat {
	b.occupants = append(b.occupants, name)
	return b
}

func (b Boat) Manifest() {
	fmt.Println("The", b.Name, "has the following occupants:")
	for _, n := range b.occupants {
		fmt.Println("\t", n)
	}
}

func main() {
	b := &Boat{
		Name: "S.S. DigitalOcean",
	}

	b.AddOccupant("Sammy the Shark")
	b.AddOccupant("Larry the Lobster")

	b.Manifest()
}