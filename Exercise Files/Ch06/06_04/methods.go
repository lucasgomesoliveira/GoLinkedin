package main

import "fmt"

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)

}

func (d Dog) SpeakThreeTimes() {
	fmt.Println(d.Sound, d.Sound, d.Sound)

}

func main() {
	poodle := Dog{"Poodle", 37, "Woof"}
	fmt.Println(poodle)
	poodle.Speak()

	poodle.Sound = "Arf"
	poodle.SpeakThreeTimes()
}
