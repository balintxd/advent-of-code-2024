package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	guard := findGuard(lines) // Y, X

    guardDir := "up"
    directions := []string{"up", "right", "down", "left"}
	movesY := []int{-1, 0, 1, 0}
	movesX := []int{0, 1, 0, -1}

	height := len(lines)
    width := len(lines[0])

    /* First part */
    visitedFields := [][2]int{}
    visitedFields = append(visitedFields, guard)

    for {
        dir := slices.Index(directions, guardDir)
        nextY := guard[0] + movesY[dir]
        nextX := guard[1] + movesX[dir]

        if nextY < 0 || nextX < 0 || nextY >= height || nextX >= width {
            break
        } else if lines[nextY][nextX] == '.' || lines[nextY][nextX] == '^' {
            guard = [2]int{nextY, nextX}
            visitedFields = addToVisitedFields(guard, visitedFields)
        } else if lines[nextY][nextX] == '#' {
            guardDir = turnRight(guardDir)
        }
    }

    fmt.Println("The guard visits",len(visitedFields),"distinct positions")
    /* End of first part */


    /* Second part */
    var newObstacle [2]int
    loop := false
    loops := 0
    var counter int

    for i := 1; i < len(visitedFields); i++ {
        guard = findGuard(lines)
        visitedFieldDirections := make(map[[3]int]bool)
        visitedFieldDirections[[3]int{guard[0], guard[1], 0}] = true
        guardDir = "up"
        newObstacle = visitedFields[i]
        counter = 0

        for {
            counter += 1
            dir := slices.Index(directions, guardDir)
            nextY := guard[0] + movesY[dir]
            nextX := guard[1] + movesX[dir]

            if (nextY < 0 || nextX < 0 || nextY >= height || nextX >= width) {
                break
            } else if (lines[nextY][nextX] == '.' || lines[nextY][nextX] == '^') && newObstacle != [2]int{nextY, nextX} {
                guard = [2]int{nextY, nextX}
                visitedFieldDirections, loop = addToVisitedFieldDirections([3]int{nextY, nextX, dir}, visitedFieldDirections)
                if loop {
                    loops += 1
                    break
                }
            } else if lines[nextY][nextX] == '#' || newObstacle == [2]int{nextY, nextX} {
                guardDir = turnRight(guardDir)
            }
        }
    }

    fmt.Println("There are",loops,"scenarios where the guard can get looped with an additional obstacle")
    /* End of second part */
}

func findGuard(lines []string) ([2]int) {
	for i, line := range lines {
		for j, char := range line {
			if char == '^' {
				return [2]int{i, j}
			}
		}
	}
	panic("Guard was not found")
}

func turnRight(currentDirection string) (string) {
    directions := []string{"up", "right", "down", "left"}
    index := slices.Index(directions, currentDirection)
    index += 1
    if (index >= len(directions)) {
        index = 0
    }
    return directions[index]
}

func addToVisitedFields(field [2]int, visitedFields [][2]int) ([][2]int) {
    for _, visitedField := range visitedFields {
        if (field == visitedField) {
            return visitedFields
        }
    }
    visitedFields = append(visitedFields, field)
    return visitedFields
}

func addToVisitedFieldDirections(field [3]int, visitedFieldDirections map[[3]int]bool) (map[[3]int]bool, bool) {
    if visitedFieldDirections[field] {
        return visitedFieldDirections, true
    }

    visitedFieldDirections[field] = true
    return visitedFieldDirections, false
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
