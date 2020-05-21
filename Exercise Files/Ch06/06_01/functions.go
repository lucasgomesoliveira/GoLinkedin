package main

import "fmt"

func main() {

	doSomething()
	sum := addValues(22, 33)
	fmt.Println(sum)

	sum = addAllValues(2, 5, 8)
	fmt.Println(sum)
}
func doSomething() {
	fmt.Println("Do Something")
}

func addValues(val1, val2 int) int {
	return val1 + val2
}

func addAllValues(vals ...int) int {
	sum := 0
	for i := range vals {

		sum += vals[i]

	}

	fmt.Printf("%T\n", vals)

	return sum
}
