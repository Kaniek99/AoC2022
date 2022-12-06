package main

import "fmt"

type Ship struct {
	Capacity                int
	Stacks                  int
	ArrangementOfContainers [][]string
}

// TODO: change way in which we store containers to keep similar graphic representation (similar to representation in input.txt)
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

func (ship Ship) getContainerAtTheTop(stack int) ( /*container*/ string /*position*/, int) {
	for i := ship.Capacity - 1; i >= 0; i-- {
		if ship.ArrangementOfContainers[i][stack] == "-" {
			continue
		} else {
			return ship.ArrangementOfContainers[i][stack], i
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

func main() {
	// reading.ReadFromFile("./day5/input.txt")
	ship := Ship{Capacity: 3, Stacks: 2}
	ship.createEmptyArrangement()
	fmt.Println(ship.ArrangementOfContainers)
	ship.addContainerToStack("A", 1)
	fmt.Println(ship.ArrangementOfContainers)
}
