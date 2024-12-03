package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	text := "";
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		text += line
	}

	/* First part */
	muls := FindMultiplications(text)
	firstSum := 0
	for _, mul := range muls {
		firstSum += CalculateMultiplication(mul)
	}

	fmt.Println("The sum of the multiplications is:", firstSum)
	/* End of first part */

	/* Second part */
	tokens := FindTokens(text)
	secondSum := 0
	process := true
	for _, token := range tokens {
		if token == "do()" {
			process = true
		} else if (token == "don't()") {
			process = false
		} else if (process) {
			secondSum += CalculateMultiplication(token)
		}
	}

	fmt.Println("The sum of the filtered multiplications is", secondSum)
	/* End of second part */
}

func FindMultiplications(text string) ([]string) {
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		panic(err)
	}

	return r.FindAllString(text, -1)
}

func FindTokens(text string) ([]string) {
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	if err != nil {
		panic(err)
	}

	return r.FindAllString(text, -1)
}

func CalculateMultiplication(mul string) (int) {
	r, err := regexp.Compile(`\d{1,3}`)
	if err != nil {
		panic(err)
	}

	numbers := r.FindAllString(mul, -1)

	first, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}
	second, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}

	return first * second
}

func ReadLines(path string) ([]string, error) {
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
