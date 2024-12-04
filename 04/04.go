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

	/* First part */
	xmasCount := CountXmas(lines)
	fmt.Println("The number of \"XMAS\" is:", xmasCount)

	/* Second part */
	crossMasCount := CountCrossMas(lines)
	fmt.Println("The number of X \"MAS\" is:", crossMasCount)
}

func CountCrossMas(lines []string) (int) {
	cols := len(lines[0])
	rows := len(lines)

	correctWords := []string{"MSAMS", "MMASS", "SSAMM", "SMASM"}

	dirX := []int{2, 1, 0, 2};
	dirY := []int{0, 1, 2, 2}

	result := 0
	word := "";

	for y := range rows {
		for x := range cols {

			word += string(lines[y][x]);

			for i := range dirX {

				moveX := x + dirX[i]
				moveY := y + dirY[i]
				
				if ValidCoordinate(moveX, moveY, cols, rows) {
					word += string(lines[moveY][moveX])
				}

			}
			if (slices.Contains(correctWords, word)) {
				result += 1
			}
			word = "";

		}
	}

	return result
}

func CountXmas(lines []string) (int) {
	cols := len(lines[0])
	rows := len(lines)

	dirX := []int{1, 1, 0, -1, -1, -1, 0, 1};
	dirY := []int{0, 1, 1, 1, 0, -1, -1, -1}

	result := 0;
	word := "";

	for y := range rows {
		for x := range cols {

			for dir := range dirX {

				word += string(lines[y][x]);

				for i := range 3 {

					moveX := x + ((i + 1) * dirX[dir])
					moveY := y + ((i + 1) * dirY[dir])
					
					if ValidCoordinate(moveX, moveY, cols, rows) {
						word += string(lines[moveY][moveX])
					}

				}
				if (word == "XMAS") {
					result += 1
				}
				word = "";
			}

		}
	}

	return result
}

func ValidCoordinate(x int, y int, cols int, rows int) (bool) {
	return x >= 0 && y >= 0 && x < cols && y < rows
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
