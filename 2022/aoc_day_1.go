package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read in input
	f, err := os.Open("./aoc_day_1_input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	topThree := make([]int, 3)
	total := 0
	for s.Scan() {
		// Insert calories carried into top3 slice while maintaining order
		if s.Text() == "" {
			temp := total
			for i, v := range topThree {
				if total > v {
					temp = v
					topThree[i] = total
					total = temp
				}
			}

			total = 0
			continue
		}

		tInt, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}

		total += tInt
	}

	// Calculate total of top3
	topTotal := 0
	for _, v := range topThree {
		topTotal += v
	}

	// Print solutions for each part
	fmt.Printf("Part One: %v\n", topThree[0])
	fmt.Printf("Part Two: %v\n", topTotal)
}
