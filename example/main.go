package main

import (
	"fmt"
)

type Sleeper interface {
	Sleep()
}

type Eater interface {
	Eat(foodName string)
}

type LazyAnimal interface {
	Sleeper
	Eater
}

type Animal struct {
	Name string `json:"name,omitempty"  xml:"name"`
	Age  int    `json:"age,omitempty"  xml:"age"`
}

type Dog struct {
	Animal
	DogLittleName string `json:"dog_little_name,omitempty"  xml:"dog_little_name"`
}

type Cat struct {
	Animal
	CatLittleName string `json:"cat_little_name,omitempty"  xml:"cat_little_name"`
}

func (d Dog) Sleep() {
	fmt.Printf("Dog %s is sleeping...%s\n", d.Name, d.DogLittleName)
}

func (d Dog) Eat(foodName string) {
	fmt.Printf("dog %s is eating %s\n", d.Name, foodName)
}

func (c Cat) Sleep() {
	fmt.Printf("Cat %s is sleeping...%s\n", c.Name, c.CatLittleName)
}

func (c Cat) Eat(foodName string) {
	fmt.Printf("cat %s is eating %s\n", c.Name, foodName)
}

func main() {
	d := Dog{Animal: Animal{"doggg", 10}, DogLittleName: "sd"}
	c := Cat{Animal: Animal{"cattt", 15}, CatLittleName: "sc"}
	lazyAnimals := []LazyAnimal{d, c}
	for _, animal := range lazyAnimals {
		animal.Eat("anc")
		animal.Sleep()
	}
}

func test() {
	d := Dog{Animal: Animal{"doggg", 10}, DogLittleName: "sd"}
	c := Cat{Animal: Animal{"cattt", 15}, CatLittleName: "sc"}
	d.Sleep()
	d.Eat("noodles")
	fmt.Println()
	c.Sleep()
	c.Eat("beef")
}
