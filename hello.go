package main

import "fmt"

func mars_age(age int) int {
	a := 365.0 / 687.0
	fmt.Println("A:", a)
	var b int = int(a * float64(age))
	return b
}

func main() {
	var age int
	fmt.Scanln(&age)
	fmt.Print(age)
	mars := mars_age(age)
	fmt.Println(mars)

}
