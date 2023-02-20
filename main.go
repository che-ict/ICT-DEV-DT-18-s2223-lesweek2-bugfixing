package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// read lines from input.txt
// in this specific case should print 10404
func main() {
	score := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		score += GetOutcome(line[0], line[2])
	}

	fmt.Println(score)
}

// Assuming the input for move is A,B or C and counter is X,Y or Z, both correspond with rock, paper, scissors in that order
// Returns 0 for loss, 3 for a tie, 6 for win + 1, 2 or 3 for using rock, paper or scissors respectively
func GetOutcome(move, counter byte) int {
	return (int(counter-move)-22)*3 + int(counter) - 87
}
