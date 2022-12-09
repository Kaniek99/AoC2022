package main

import (
	"aoc2022/reading"
	"errors"
	"fmt"
)

type Ship struct {
	Capacity                int
	Stacks                  int
	ArrangementOfContainers [][]string
}

// TODO: refactor this, implement stack
func (ship *Ship) createEmptyArrangement() {
	ship.ArrangementOfContainers = make([][]string, ship.Capacity)
	for i := range ship.ArrangementOfContainers {
		ship.ArrangementOfContainers[i] = make([]string, ship.Stacks)
	}
	for i := 0; i < ship.Capacity; i++ {
		for j := 0; j < ship.Stacks; j++ {
			ship.ArrangementOfContainers[i][j] = "-"
		}
	}
}

func (ship Ship) getContainerAtTheTop(stack int) (string, int) {
	for i := ship.Capacity - 1; i >= 0; i-- {
		if ship.ArrangementOfContainers[i][stack] == "-" {
			continue
		} else {
			return ship.ArrangementOfContainers[i][stack], i // return container, position
		}
	}
	return "-", 0
}

func (ship *Ship) removeContainerFromStack(stack int) string {
	removedContainer := ""
	for i := ship.Capacity - 1; i >= 0; i-- {
		if ship.ArrangementOfContainers[stack][i] == "-" {
			continue
		} else {
			removedContainer = ship.ArrangementOfContainers[stack][i]
			ship.ArrangementOfContainers[stack][i] = "-"
			return removedContainer
		}
	}
	return "There are no containers on this stack"
}

func (ship *Ship) addContainerToStack(container string, stack int) bool {
	topContainer, peak := ship.getContainerAtTheTop(stack)
	if topContainer == "-" {
		ship.ArrangementOfContainers[0][stack] = container
		return true
	} else {
		if peak+1 < ship.Capacity {
			ship.ArrangementOfContainers[peak+1][stack] = container
			return true
		} else {
			//error handling
			return false
		}
	}
}

func (ship *Ship) changeCapacity(capacity int) {
	ship.Capacity = capacity
}

func separateInputAndInstruction(data []string) (int, error) {
	for i, elem := range data {
		if elem == "" {
			return i, nil
		}
	}
	return 0, errors.New("invalid data format")
}

func main() {
	data := reading.ReadFromFile("./day5/input.txt")
	ship := Ship{Capacity: 3, Stacks: 2}
	ship.createEmptyArrangement()
	// fmt.Println(ship.ArrangementOfContainers)
	// ship.addContainerToStack("A", 1)
	// fmt.Println(ship.ArrangementOfContainers)
	// ship.addContainerToStack("B", 1)
	// fmt.Println(ship.ArrangementOfContainers)
	separator, err := separateInputAndInstruction(data[:3]) // separator is an index of blank line in input.txt
	fmt.Println(separator)
	if err != nil {
		fmt.Println(err)
		return
	}

}
