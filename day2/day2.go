package main

import (
	"aoc2022/reading"
	"fmt"
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
	data := reading.ReadFromFile("./day2/input.txt")

	myPoints := 0
	elfPoints := 0
	for _, line := range data {
		battle := Battle{elf: int(line[0]) - 64, me: int(line[2]) - 87}
		switch battle.result() {
		// score for a single round is the score for the shape you selected (1 for Rock, 2 for Paper, and 3 for Scissors)
		// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won)
		case Draw:
			elfPoints += battle.elf + 3
			myPoints += battle.me + 3
		case Lose:
			elfPoints += battle.elf + 6
			myPoints += battle.me
		case Win:
			elfPoints += battle.elf
			myPoints += battle.me + 6
		}
	}
	fmt.Println("Points earned by me:", myPoints)
	fmt.Println("Points earned by elf:", elfPoints)

	//second part: X means I need to lose, Y draw, Z win X <-> 1 Y <-> 2 Z <-> 3
	myPointsPart2 := 0
	elfPointsPart2 := 0
	for _, line := range data {
		battle := Battle{elf: int(line[0]) - 64, me: int(line[2]) - 87}
		switch battle.me {
		case 1:
			if battle.elf == 1 {
				elfPointsPart2 += battle.elf + 6
				myPointsPart2 += 3
			} else {
				elfPointsPart2 += battle.elf + 6
				myPointsPart2 += battle.elf - 1
			}
		case 2:
			elfPointsPart2 += battle.elf + 3
			myPointsPart2 += battle.elf + 3
		case 3:
			if battle.elf == 3 {
				elfPointsPart2 += battle.elf
				myPointsPart2 += 7
			} else {
				elfPointsPart2 += battle.elf
				myPointsPart2 += battle.elf + 1 + 6
			}
		}
	}
	fmt.Println("Points earned by me:", myPointsPart2)
	fmt.Println("Points earned by elf:", elfPointsPart2)
}
