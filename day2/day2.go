package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Draw = iota
	Lose
	Win
)

type Battle struct {
	elf int
	me  int
}

func (battle Battle) result() int {
	return ((battle.elf-battle.me)%3 + 3) % 3
}

func main() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	data := []string{}

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	file.Close()

	myPoints := 0
	elfPoints := 0
	for _, line := range data {
		// rock = 0 paper = 1 scisor = 2
		battle := Battle{elf: int(line[0]) - 65, me: int(line[2]) - 88}
		switch battle.result() {
		// score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
		// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
		case Draw:
			elfPoints += battle.elf + 1 + 3
			myPoints += battle.me + 1 + 3
		case Lose:
			elfPoints += battle.elf + 1 + 6
			myPoints += battle.me + 1
		case Win:
			elfPoints += battle.elf + 1
			myPoints += battle.me + 1 + 6
		}
	}
	fmt.Println("Points earned by me:", myPoints)
	fmt.Println("Points earned by elf:", elfPoints)

	//second part X means I need to lose, Y draw, Z win X <-> 0 Y <-> 1 Z <-> 2
	myPointsPart2 := 0
	elfPointsPart2 := 0
	for _, line := range data {
		battle := Battle{elf: int(line[0]) - 65, me: int(line[2]) - 88}
		switch battle.me {
		case 0:
			elfPointsPart2 += battle.elf + 1 + 6
			myPointsPart2 += (battle.elf+1)%3 + 1
		case 1:
			elfPointsPart2 += battle.elf + 1 + 3
			myPointsPart2 += battle.elf + 1 + 3
		case 2:
			elfPointsPart2 += battle.elf + 1
			myPointsPart2 += (battle.elf-1)%3 + 6
		}
	}
	fmt.Println("Points earned by me:", myPointsPart2)
	fmt.Println("Points earned by elf:", elfPointsPart2)
}

//wrong answers: 11225 12083 12266
