package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	Result int
	Operands []int
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	equations := GetEquations(lines)

	/* First part */
	sumBase2 := 0
	for _, equation := range equations {
		if equation.CorrectEquation(2) {
			sumBase2 += equation.Result
		}
	}
	fmt.Println(sumBase2)
	/* End of first part */

	/* Second part */
	sumBase3 := 0
	for _, equation := range equations {
		if equation.CorrectEquation(3) {
			sumBase3 += equation.Result
		}
	}
	fmt.Println(sumBase3)
	/* End of second part */
}

// part of the logic is from https://github.com/racecraftr/advent-of-code/blob/main/y2024/day07/main.go
func (equation *Equation) CorrectEquation(base int) (bool) {
	operators := [3]rune{'+', '*', '|'}
	permCount := int(math.Pow(float64(base), float64(len(equation.Operands) - 1)))

	outer:
	for permNum := range permCount {
		perm, result := permNum, equation.Operands[0]
		for i, next := range equation.Operands[1:] {
			switch operators[perm % base] {
				case '+':
					result += next
				case '*':
					result *= next
				case '|':
					leftNum := strconv.Itoa(result)
					rightNum := strconv.Itoa(next)
					concatenatedNum := leftNum + rightNum
					result, _ = strconv.Atoi(concatenatedNum)
			}

			if result == equation.Result && i == len(equation.Operands[1:]) - 1 {
				return true
			}
			if result > equation.Result {
				continue outer
			}
			perm /= base
		}
	}
	return false
}

func GetEquations(lines []string) ([]Equation) {
	equations := []Equation{}

	for _, line := range lines {
		parts := strings.Split(line, ":")

		// Result
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		// Operands
		operands := []int{}
		stringOperands := strings.Split(parts[1][1:], " ")
		for _, stringOperand := range stringOperands {
			operand, err := strconv.Atoi(stringOperand)
			if err != nil {
				panic(err)
			}
			operands = append(operands, operand)
		}

		equations = append(equations, Equation{
			Result: result,
			Operands: operands,
		})
	}

	return equations
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
