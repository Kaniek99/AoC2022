package main

import (
	"aoc2022/reading"
	"fmt"
	"strings"
)

func checkRepeating(contentOfRucksack string) int {
	for _, item := range contentOfRucksack[:len(contentOfRucksack)/2] {
		if strings.ContainsRune(contentOfRucksack[len(contentOfRucksack)/2:], item) {
			if item >= 97 && item <= 122 {
				return int(item - 96) // a <-> 1
			} else {
				return int(item - 38) // A <-> 27
			}
		}
	}
	return 0
}

func checkGroup(firstRucksack, secondRucksack, thirdRucksack string) int {
	for _, item := range firstRucksack {
		if !strings.ContainsRune(secondRucksack, item) {
			continue
		} else {
			if !strings.ContainsRune(thirdRucksack, item) {
				continue
			} else {
				if item >= 97 && item <= 122 {
					return int(item - 96) // a <-> 1
				} else {
					return int(item - 38) // A <-> 27
				}
			}
		}
	}
	return 0
}

func main() {
	data := reading.ReadFromFile("./day3/input.txt")

	sum := 0
	for _, elem := range data {
		sum += checkRepeating(elem)
	}
	fmt.Println("sum:", sum)

	sumForBadges := 0
	for i := 0; i < len(data); i += 3 {
		sumForBadges += checkGroup(data[i], data[i+1], data[i+2])
	}
	fmt.Println("sum:", sumForBadges)
}
