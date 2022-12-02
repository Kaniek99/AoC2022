package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./day1/input.txt")
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
