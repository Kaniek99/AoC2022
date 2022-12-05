package main

import (
	"aoc2022/reading"
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	data := reading.ReadFromFile("./day1/input.txt")

	sum := 0
	sumCalories := []int{}
	for _, elem := range data {
		if elem == "" {
			sumCalories = append(sumCalories, sum)
			sum = 0
		} else {
			calories, err := strconv.Atoi(elem)
			if err != nil {
				log.Fatal(err)
			} else {
				sum += calories
			}
		}
	}
	sort.Ints(sumCalories)
	fmt.Println("highest number of calories carried by an elf:", sumCalories[len(sumCalories)-1])
	fmt.Println("sum of calories carried by 3 most loaded elves:", sumCalories[len(sumCalories)-1]+sumCalories[len(sumCalories)-2]+sumCalories[len(sumCalories)-3])
}
