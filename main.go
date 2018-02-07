/*
  -- PANCAKE FLIPPER COUNTER --
  for Infinite Pancake House.
  This program will properly determine the minimum number of
  flips required to properly orientate a randomly orientated
  stack of Happy Face Pancakes happy side up.
  Analyzing each stack using X-ray vision, test cases have been
  established. Cakes requiring a flip are denoted with a 1
  (indicating true, or yes, it needs a flip) cakes happy side up
  not requiring a flip are denoted with a 0 in the array
  (indicating false, or no, it does not need to be flipped).

	from the command line run:
	go run main.go
*/

package main

import (
	"fmt"
)

func main() {

	cases := [] []int{ // test cases
		{1},
		{1, 0},
		{0, 1},
		{0, 0, 0},
		{1, 1, 0, 1},
	}

	fmt.Println("*** Pancake Flipper ***\n")
	fmt.Println("running given test cases:")

	for i := range cases { // run all test cases
		fmt.Println(descp(i + 1), orientatePancakes(cases[i]))
	}

	//simpleFunctionTests()
}

// returns true if all pancakes are orientated the same direction.
func allSame(s []int) bool {
	result := true
	which := s[0]

	for i := range s {
		if s[i] != which {
			result = false
		}
	}

	return result
}

// returns true if all pancakes are smile side up.
func allSmileUp(s []int) bool {
	return allSame(s) && s[0] == 0
}

func findDiff(s []int, val int) int {
	for i := range s {
		if s[i] != val {
			return i
		}
	}
	return 0
}

// flips a range of pancakes from 0 to 1 and vice versa
func flip(s []int) []int {
	for i := range s {
		if s[i] == 1 {
			s[i] = 0
		} else {
			s[i] = 1
		}
	}
	return s
}

// takes a set of matching pancakes and flips them to match the next consecutive pancake outside the set
func flipPancakes(s []int) []int {
	toMatch := s[0]
	end := findDiff(s, toMatch)

	flip(s[0: end])

	return s
}

// helper method to fill a slice with the same value (other than 0, of course)
func fill(s []int, val int) []int {
	for i := range s {
		s[i] = val
	}
	return s
}

// recursive call to orientate all pancakes.
func orientatePancakes(s []int) int {
	if allSame(s) {
		if allSmileUp(s) {
			return 0
		} else {
			return 1
		}
	} else {
		n := orientatePancakes(flipPancakes(s))
		return n + 1
	}
}

// helper function to describe label test cases
func descp(i int) string {
	return fmt.Sprintf("Case #%d:", i)
}


// function tests
func simpleFunctionTests() {
	smileUp := make([]int, 5)
	blankUp := fill(make([]int, 5), 1)
	mixed := []int {1,1,0,0,1}

	fmt.Println("\nrunning function tests:")
	fmt.Println("allSame smile up:", allSame(smileUp))
	fmt.Println("allSame blank up:", allSame(blankUp))
	fmt.Println("findDiff:", findDiff(mixed, 0))
	fmt.Println("allSmileUp blankUp:", allSmileUp(blankUp))
	fmt.Println("allSmileUp smileUp:", allSmileUp(smileUp))
	fmt.Println("flip:", flip(blankUp))

	flipPancakes(mixed)
	fmt.Println(mixed)
}
