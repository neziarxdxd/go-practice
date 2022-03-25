package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	fmt.Scan(&num3)
	toNname(num1)
	toNname(num2)
	toNname(num3)
}

func toNname(num int) {
	if num == 0 {
		fmt.Println("Zero")
	} else if num == 1 {
		fmt.Println("One")
	} else if num == 2 {
		fmt.Println("Two")
	} else if num == 3 {
		fmt.Println("Three")
	} else if num == 4 {
		fmt.Println("Four")
	} else if num == 5 {
		fmt.Println("Five")
	} else if num == 6 {
		fmt.Println("Six")
	} else if num == 7 {
		fmt.Println("Seven")
	} else if num == 8 {
		fmt.Println("Eight")
	} else if num == 9 {
		fmt.Println("Nine")
	} else if num == 10 {
		fmt.Println("Ten")
	}

}
