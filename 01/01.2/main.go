package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	mapB := make(map[int]int)
	readInput(inputFile, listA, mapB)

}

func readInput(inputFile *os.File, listA []int, mapB map[int]int) {
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) < 2 {
			log.Printf("Bad input line: %+v\n", line)
			continue
		}
		addStrToArray(line[0], &listA)
		addStrToMap(line[1], mapB)
	}
	fmt.Printf("listA: %+v\n", listA)
	fmt.Printf("mapB: %+v\n", mapB)
	simScore := calculateSimilarity(listA, mapB)
	fmt.Printf("Similarity Score: %d\n", simScore)
}

func calculateSimilarity(listA []int, mapB map[int]int) int {
	simScore := 0
	for _, val := range listA {
		if mapVal, exists := mapB[val]; exists {
			simScore += val * mapVal
		}
	}

	return simScore
}

func addStrToMap(num string, mapInts map[int]int) {
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalf("Unable to convert string %s to int: %s", intNum, err)
	}
	if val, exists := mapInts[intNum]; exists {
		mapInts[intNum] = val + 1
	} else {
		mapInts[intNum] = 1
	}
}

func addStrToArray(num string, listInts *[]int) {
	intNum, err := strconv.Atoi(num)
	if err != nil {
		log.Fatalf("Unable to convert string %s to int: %s", intNum, err)
	}
	*listInts = append(*listInts, intNum)
}
