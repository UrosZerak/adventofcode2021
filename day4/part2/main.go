package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	column []int
	row    []int
}

type Solution struct {
	index            int
	RowOrColumn      string
	lastCalledNumber int
}

func main() {
	file, err := os.Open("../input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	selectedNumbers := strings.Split(scanner.Text(), ",")

	positionsMap := make(map[int]Position)
	solutionsMapRow := make(map[int]int)
	solutionsMapColumn := make(map[int]int)
	solution := []Solution{}
	indexOfRow := 0
	indexOfColumn := 0

	var bingoBoards [][]int

	for scanner.Scan() {
		if len(scanner.Text()) != 0 {

			stringNumbers := strings.Split(scanner.Text(), " ")
			numbers := []int{}

			for i := 0; i < len(stringNumbers); i++ {
				a, err := strconv.Atoi(stringNumbers[i])
				if err == nil {
					numbers = append(numbers, a)
				}
			}
			for i := 0; i < len(numbers); i++ {
				column := positionsMap[numbers[i]].column
				row := positionsMap[numbers[i]].row
				column = append(column, i+indexOfColumn)
				row = append(row, indexOfRow)
				position := Position{column, row}
				positionsMap[numbers[i]] = position
			}
			bingoBoards = append(bingoBoards, numbers)
			indexOfRow++
		} else {
			indexOfColumn += 5
		}
	}

	for _, x := range selectedNumbers {
		numberX, _ := strconv.Atoi(x)
		if len(positionsMap[numberX].column) > 0 {
			for _, v := range positionsMap[numberX].column {
				solutionsMapColumn[v]++
				if solutionsMapColumn[v] == 5 {
					lastCalledNumber := numberX
					solution = append(solution, Solution{index: v, RowOrColumn: "column", lastCalledNumber: lastCalledNumber})
				}
			}
		}

		if len(positionsMap[numberX].row) > 0 {
			for _, v := range positionsMap[numberX].row {
				solutionsMapRow[v]++
				if solutionsMapRow[v] == 5 {
					lastCalledNumber := numberX
					solution = append(solution, Solution{index: v, RowOrColumn: "row", lastCalledNumber: lastCalledNumber})
				}
			}
		}

	}

	var markedNumbers = make(map[int]bool)
	for i := 0; i < len(selectedNumbers); i++ {
		n, _ := strconv.Atoi(selectedNumbers[i])
		markedNumbers[n] = true
		if selectedNumbers[i] == strconv.Itoa(solution[len(solution)-1].lastCalledNumber) {
			break
		}
	}

	sumOfUnmarkedNumbers := 0
	columnToRow := 0
	if solution[len(solution)-1].RowOrColumn == "column" {
		columnToRow = 5
	}

	var bingoBoard int = solution[len(solution)-1].index - solution[len(solution)-1].index%5 - columnToRow
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !markedNumbers[bingoBoards[bingoBoard][j]] {
				sumOfUnmarkedNumbers += bingoBoards[bingoBoard][j]
			}
		}
		bingoBoard++
	}
	fmt.Println(solution[len(solution)-1].lastCalledNumber * sumOfUnmarkedNumbers)
	checkError(scanner.Err())

}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// idea make a 2d array to save bingo boards, linked list of neighbors 18 <->99<->39
//check for depth of substring as well for
