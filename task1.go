package main

import "fmt"

func fillJugs(size1, size2, target int) {
	var jug1, jug2 int

	fmt.Printf("Initial State: jug1 = %d gallons, jug2 = %d gallons\n", jug1, jug2)

	for jug1 != target && jug2 != target {
		if jug1 == 0 {
			jug1 = size1
			fmt.Printf("Fill the first jug: jug1 = %d gallons\n", jug1)
		} else if jug2 == size2 {
			jug2 = 0
			fmt.Printf("Empty the second jug: jug2 = %d gallons\n", jug2)
		} else {
			pour := min(jug1, size2-jug2)
			jug1 -= pour
			jug2 += pour
			fmt.Printf("Pour water from the first jug to the second jug: jug1 = %d gallons, jug2 = %d gallons\n", jug1, jug2)
		}
	}

	fmt.Printf("Final State: jug1 = %d gallons, jug2 = %d gallons\n", jug1, jug2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fillJugs(4, 3, 2)
}
