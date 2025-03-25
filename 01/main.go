package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFileName := "./input.txt"

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Error: unable to open input file: %s", err)
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Printf("Unable to close inputFile: %s\n", err)
		}
	}(inputFile)

	listA := make([]int, 0)
	listB := make([]int, 0)

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) < 2 {
			log.Printf("Bad input line: %+v\n", line)
			continue
		}
		addStrToArray(line[0], &listA)
		addStrToArray(line[1], &listB)
	}

	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})
	sort.Slice(listB, func(i, j int) bool {
		return listB[i] < listB[j]
	})

	log.Printf("listA %+v", listA)
	log.Printf("listB %+v", listB)

	distance := computeDistance(listA, listB)
	log.Printf("Total Distance: %d", distance)
}

func addStrToArray(num string, listInts *[]int) {
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalf("Unable to convert string %s to int: %s", intNum, err)
	}
	*listInts = append(*listInts, intNum)
}

func sortCondition(a, b int) bool {
	return a < b
}

func computeDistance(listA, listB []int) int {
	if len(listA) != len(listB) {
		log.Fatalf("Length of the 2 inputs are not equal!\n\n%+v\n\n%+v", listA, listB)
	}
	totalDistance := 0
	for i := range len(listA) {
		totalDistance += absInt(listA[i], listB[i])
	}
	return totalDistance
}

func absInt(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
