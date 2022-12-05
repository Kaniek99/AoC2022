package reading

import (
	"bufio"
	"log"
	"os"
)

func ReadFromFile(pathToFile string) []string {
	file, err := os.Open(pathToFile)
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
	return data
}
