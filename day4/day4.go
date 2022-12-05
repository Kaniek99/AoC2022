package main

import (
	"aoc2022/reading"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func hasAll(firstAssignment, secondAssignment []string) bool {
	beginOfFirst, err := strconv.Atoi(firstAssignment[0])
	endOfFirst, err := strconv.Atoi(firstAssignment[1])
	beginOfSecond, err := strconv.Atoi(secondAssignment[0])
	endOfSecond, err := strconv.Atoi(secondAssignment[1])
	if err != nil { // error handling should be improved
		log.Fatal(err)
	} else {
		if (beginOfFirst >= beginOfSecond && endOfFirst <= endOfSecond) || (beginOfSecond >= beginOfFirst && endOfSecond <= endOfFirst) {
			return true
		}
	}
	return false
}

func hasAny(firstAssignment, secondAssignment []string) bool {
	beginOfFirst, err := strconv.Atoi(firstAssignment[0])
	endOfFirst, err := strconv.Atoi(firstAssignment[1])
	beginOfSecond, err := strconv.Atoi(secondAssignment[0])
	endOfSecond, err := strconv.Atoi(secondAssignment[1])
	if err != nil { // error handling should be improved
		log.Fatal(err)
	} else {
		if math.Max(float64(beginOfFirst), float64(beginOfSecond)) <= math.Min(float64(endOfFirst), float64(endOfSecond)) {
			return true
		}
	}
	return false
}

func main() {
	data := reading.ReadFromFile("./day4/input.txt")
	counter := 0
	counterForAny := 0
	for _, elem := range data {
		assignmentForPair := strings.Split(elem, ",")
		firstAssignment := strings.Split(assignmentForPair[0], "-")
		secondAssignment := strings.Split(assignmentForPair[1], "-")
		if hasAll(firstAssignment, secondAssignment) {
			counter += 1
		}
		if hasAny(firstAssignment, secondAssignment) {
			counterForAny += 1
		}
	}
	fmt.Println("overlaps:", counter)
	fmt.Println("assignments with atleast one same section:", counterForAny)
}
