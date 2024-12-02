package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	safeCount := safeReportCount(lines, false)
	safeCountProblemDapened := safeReportCount(lines, true)

	/* First part */
	fmt.Println("The number of safe reports is:", safeCount)

	/* Second part */
	fmt.Println("The numebr of safe reports using the problem dampener is:", safeCountProblemDapened)
}

func safeReportCount(lines []string, problemDapener bool) (int) {
	safeCount := 0

	for _, line := range lines {
		elements := strings.Split(line, " ")
		var numbers = []int{}

		for _, element := range elements {
			number, err := strconv.Atoi(element)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
		}

		if (problemDapener) {
			if (safeNumbers(numbers, false) || safeNumbers(numbers, true)) {
				safeCount += 1
			} else {
				for i := range numbers {
					newNumbers := removeIndex(numbers, i)
					if (safeNumbers(newNumbers, false) || safeNumbers(newNumbers, true)) {
						safeCount += 1
						break
					}
				}
			}
		} else {
			if safeNumbers(numbers, false) || safeNumbers(numbers, true) {
				safeCount += 1
			}
		}
	}

	return safeCount
}

func safeNumbers(numbers []int, increasing bool) (bool) {
	prev := -1
	
	for _, number := range numbers {
		if prev == -1 {
			prev = number
			continue
		}
		diff := int(math.Abs(float64(prev - number)))
		if diff == 0 || diff > 3 {
			return false
		}
		if increasing && number < prev {
			return false
		}
		if!increasing && number > prev {
			return false
		}
		prev = number
	}

	return true
}

func removeIndex(numbers []int, index int) []int {
	new := make([]int, 0)
	new = append(new, numbers[:index]...)
	return append(new, numbers[index+1:]...)
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
