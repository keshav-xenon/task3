package main

import (
	"fmt"
	"strconv"
)

type State struct {
	leftMissionaries  int
	leftCannibals     int
	rightMissionaries int
	rightCannibals    int
	boatIsLeft        bool
	path              []Move
}

type Move struct {
	missionaries int
	cannibals    int
}

var moves = []Move{
	{1, 0}, // One missionary crosses
	{2, 0}, // Two missionaries cross
	{0, 1}, // One cannibal crosses
	{0, 2}, // Two cannibals cross
	{1, 1}, // One missionary and one cannibal cross
}

func (s State) String() string {
	boatPosition := 0
	if s.boatIsLeft {
		boatPosition = 1
	}
	return strconv.Itoa(s.leftMissionaries) + strconv.Itoa(s.leftCannibals) + strconv.Itoa(s.rightMissionaries) + strconv.Itoa(s.rightCannibals) + strconv.Itoa(boatPosition)
}

func isValidState(state State) bool {
	return (state.leftMissionaries == 0 || state.leftMissionaries >= state.leftCannibals) &&
		(state.rightMissionaries == 0 || state.rightMissionaries >= state.rightCannibals) &&
		state.leftMissionaries >= 0 && state.leftCannibals >= 0 &&
		state.rightMissionaries >= 0 && state.rightCannibals >= 0
}

func printPath(path []Move) {
	for _, move := range path {
		fmt.Printf("Move: %dM, %dC\n", move.missionaries, move.cannibals)
	}
}

func solveMissionariesAndCannibals(initialState State) {
	queue := []State{initialState}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.leftMissionaries == 0 && current.leftCannibals == 0 {
			fmt.Println("Solution found!")
			printPath(current.path)
			return
		}

		for _, move := range moves {
			newState := current
			if current.boatIsLeft {
				newState.leftMissionaries -= move.missionaries
				newState.leftCannibals -= move.cannibals
				newState.rightMissionaries += move.missionaries
				newState.rightCannibals += move.cannibals
			} else {
				newState.rightMissionaries -= move.missionaries
				newState.rightCannibals -= move.cannibals
				newState.leftMissionaries += move.missionaries
				newState.leftCannibals += move.cannibals
			}
			newState.boatIsLeft = !current.boatIsLeft
			newState.path = append(current.path, move)

			if isValidState(newState) && !visited[newState.String()] {
				visited[newState.String()] = true
				queue = append(queue, newState)
			}
		}
	}

	fmt.Println("No solution found.")
}

func main() {
	initialState := State{
		leftMissionaries:  3,
		leftCannibals:     3,
		rightMissionaries: 0,
		rightCannibals:    0,
		boatIsLeft:        true,
		path:              []Move{},
	}

	fmt.Println("Solving the Missionaries and Cannibals problem:")
	solveMissionariesAndCannibals(initialState)
}
