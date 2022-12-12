package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part One: %v\n", partOne())
	fmt.Printf("Part Two: %v\n", partTwo())
}

func partOne() int {
	// Read inputs
	f, err := os.Open("./aoc_day_2_input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	// Score map
	//   A, X = Rock = 1 pt
	//   B, Y = Paper = 2 pts
	//   C, Z = Scissors = 3 pts
	values := map[string]int{"A": 1, "B": 2, "C": 3, "X": 1, "Y": 2, "Z": 3}

	score := 0
	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		opp, me := values[l[0]], values[l[1]]

		score += me

		// Add 6 points if you win Rock, Paper, Scissors
		// Add 3 points for a tie
		// No points for a loss
		if opp == me {
			score += 3
		} else if me == 1 && opp == 3 {
			score += 6
		} else if me > opp && !(opp == 1 && me == 3) {
			score += 6
		}
	}

	return score
}

func partTwo() int {
	// Read inputs
	f, err := os.Open("./aoc_day_2_input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	// Score map
	//   A = Rock = 1 pt
	//   B = Paper = 2 pts
	//   C = Scissors = 3 pts
	oppValues := map[string]int{"A": 1, "B": 2, "C": 3}
	winValues := map[string]int{"A": 2, "B": 3, "C": 1}
	loseValues := map[string]int{"A": 3, "B": 1, "C": 2}

	score := 0
	for s.Scan() {
		l := strings.Split(s.Text(), " ")
		opp, me := l[0], l[1]

		// Add 6 points if you win Rock, Paper, Scissors (Z)
		// Add 3 points for a tie (Y)
		// No points for a loss (X)
		if me == "Y" {
			score += 3
			score += oppValues[opp]
		} else if me == "Z" {
			score += 6
			score += winValues[opp]
		} else {
			score += loseValues[opp]
		}
	}

	return score
}
