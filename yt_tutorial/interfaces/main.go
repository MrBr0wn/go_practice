package main

import (
	"fmt"
	"interfaces/pets"
	"os"
)

func main() {
	cat := pets.Cat{
		Animal: pets.Animal{Name: "mr. buttons"},
		Age:    5,
	}
	dog := pets.Dog{
		Animal:   pets.Animal{Name: "Frank"},
		Age:      8,
		Weight:   5,
		IsAsleep: true,
	}

	var feedToCat uint8 = 3
	var feedToDog uint8 = 22

	catFeed, err := feed(&cat, feedToCat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding cat: %v\n", err)
	} else {
		fmt.Println("Cat ate:", catFeed)
	}

	fmt.Printf("\n\n------------------------\n\n")

	dogFeed, err := feed(&dog, feedToDog)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error feeding dog: %v\n", err)
	} else {
		println("Dog ate:", dogFeed)
	}

	fmt.Print("\n\n--------------------------\n\n")

	displayInfo(cat)
	displayInfo(dog)
	displayInfo(3.14)
}

func feed(animal pets.EaterWalker, amount uint8) (uint8, error) {
	println("First, let's walk!")
	println(animal.Walk())

	println("Now, let's feed our", animal.GetName())

	return animal.Eat(amount)
}

func displayInfo(i interface{}) {
	switch v := i.(type) {
	case string:
		println("This is a string:", v)
	case int:
		println("This is an int", v)
	case pets.Cat:
		fmt.Printf("This is a Cat. Name: %s, Age: %d\n", v.Name, v.Age)
	case pets.Dog:
		fmt.Printf("This is a Dog. Name: %s, Age: %d, Weight: %d\n", v.Name, v.Age, v.Weight)
	default:
		fmt.Println("Unknown type :(")
	}
}
