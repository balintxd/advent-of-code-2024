package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	var left []int
	var right []int

	for _, line := range lines {
		numbers := strings.Split(line, "   ")

		leftValue, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		rightValue, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	/* First part */
	totalDistance := totalDistance(left, right)
	fmt.Println("The total distance is:", totalDistance)

	/* Second part */
	similarityScore := similarityScore(left, right)
	fmt.Println("The similarity score is:", similarityScore)
}

func totalDistance(left []int, right []int) (int) {
	slices.Sort(left)
	slices.Sort(right)

	totalDistance := 0

	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDistance
}

func similarityScore(left []int, right []int) (int) {
	similarityScore := 0
	dict := make(map[int]int)
	for _, num := range right {
		dict[num] = dict[num] + 1
	}
	for _, num := range left {
		similarityScore = similarityScore + (num * dict[num])
	}
	return similarityScore
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}