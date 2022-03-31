package main

import "fmt"

/**
You are making a program to analyze sport match results and calculate the points of the given team.
The match results are stored in an array called results.
Each match has one of the following results:
"w" - won
"l" - lost
"d" - draw

A win adds three points, a draw adds one point, and a lost match does not add any points.

Your program needs to take the last match result as input and append it to the results array. After that, calculate and output the total points the team gained from the results.
**/
func main() {
	totalPoints := 0
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var lastResult string
	fmt.Scanln(&lastResult)
	results = append(results, lastResult)
	for _, result := range results {
		if result == "w" {
			totalPoints += 3
		} else if result == "d" {
			totalPoints += 1
		}
	}

	fmt.Println(totalPoints)

}
