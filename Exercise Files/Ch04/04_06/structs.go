package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
}

func main() {

	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v", poodle)
	fmt.Printf("\nBreed %v \nWeigth %v", poodle.Breed, poodle.Weight)

}
