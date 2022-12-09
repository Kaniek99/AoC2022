package main

import (
	"aoc2022/reading"
	"fmt"
)

func isAnyLetterRepeating(word string) bool {
	if len(word) == 1 {
		return false
	}
	for i := 1; i < len(word); i++ {
		if word[0] == word[i] {
			return true
		}
	}
	return isAnyLetterRepeating(word[1:])
}

func findGroupOfDifferentLetters(word string, sizeOfGroup int) int {
	for i := 0; i < len(word)-sizeOfGroup; i++ {
		if !isAnyLetterRepeating(word[i : i+sizeOfGroup]) {
			return i + sizeOfGroup
		}
	}
	return len(word)
}

func main() {
	data := reading.ReadFromFile("./day6/input.txt")
	numberOfProcessedLettersForPacket := findGroupOfDifferentLetters(data[0], 4)
	numberOfProcessedLettersForMessage := findGroupOfDifferentLetters(data[0], 14)

	if len(data[0]) != numberOfProcessedLettersForPacket {
		fmt.Println(numberOfProcessedLettersForPacket)
	} else {
		fmt.Println("Invalid input.")
	}

	if len(data[0]) != numberOfProcessedLettersForMessage {
		fmt.Println(numberOfProcessedLettersForMessage)
	} else {
		fmt.Println("Invalid input.")
	}
}
