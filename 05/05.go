package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := ReadLines("input.txt");
	if err != nil {
		panic(err)
	}

	rules := retrieveRules(lines)
	updates := retieveUpdates(lines)
	fmt.Println("There are",len(rules),"rules")
	fmt.Println("There are",len(updates),"updates")

	correctUpdates := [][]int{}
	incorrectUpdates := [][]int{}
	for _, update := range updates {
		if (correctUpdate(update, rules)) {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}
	fmt.Println("There are",len(correctUpdates),"correct updates")
	fmt.Println("There are",len(incorrectUpdates),"incorrect updates")

	/* First part */
	sumOfCorrectUpdateMiddles := 0
	for _, update := range correctUpdates {
		sumOfCorrectUpdateMiddles += update[len(update)/2]
	}
	fmt.Println("The sum of the middle numbers of the correct updates is:", sumOfCorrectUpdateMiddles)
	
	/* Second part */
	sumOfIncorrectUpdateMiddles := 0
	for _, update := range incorrectUpdates {
		correctedUpdate := fixUpdate(update, rules)
		sumOfIncorrectUpdateMiddles += correctedUpdate[len(correctedUpdate)/2]
	}
	fmt.Println("The sum of the middle numbers of the corrected updates is:", sumOfIncorrectUpdateMiddles)
}

func retrieveRules(lines []string) ([][]int) {
	rules := [][]int{}	
	for _, line := range lines {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			firstPart, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			secondPart, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, []int{firstPart, secondPart})
		} else {
			break;
		}
	}
	return rules
}

func retieveUpdates(lines []string) ([][]int) {
	updates := [][]int{}
	for _, line := range lines {
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			updateRow := []int{}
			for _, part := range parts {
				convertedPart, err := strconv.Atoi(part)
				if err != nil {
					panic(err)
				}
				updateRow = append(updateRow, convertedPart)
			}
			updates = append(updates, updateRow)
		}
	}
	return updates
}

func correctUpdate(update []int, rules [][]int) (bool) {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			for _, rule := range rules {
				if (rule[0] == update[j] && rule[1] == update[i]) {
					return false
				}
			}
		}
	}
	return true
}

func fixUpdate(update []int, rules [][]int) ([]int) {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			for _, rule := range rules {
				if (rule[0] == update[j] && rule[1] == update[i]) {
					newUpdate := swapIndexes(update, i, j)
					return fixUpdate(newUpdate, rules)
				}
			}
		}
	}
	return update
}

func swapIndexes(update []int, firstIndex int, secondIndex int) ([]int) {
	temp := update[firstIndex]
	update[firstIndex] = update[secondIndex]
	update[secondIndex] = temp
	return update
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