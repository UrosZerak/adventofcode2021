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
	row []int
}

type Solution struct {
	index int
	RowOrColumn string
	lastCalledNumber int
}

func main() {
	file, err := os.Open("./input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	selectedNumbers := strings.Split(scanner.Text(),",")

	positionsMap := make(map[int]Position)
	solutionsMapRow := make(map[int]int)
	solutionsMapColumn := make(map[int]int)
	solution := Solution{index: -1}
	indexOfRow := 0
	indexOfColumn := 0

	var bingoBoards [][5]int

	for scanner.Scan(){
		if len(scanner.Text()) != 0{

			stringNumbers := strings.Split(scanner.Text()," ")
			var numbers [5]int

			separtedArrayString := strings.Join(stringNumbers,",")

			stringNumberWithoutSpace := strings.ReplaceAll(separtedArrayString, ",,",",")
			if(stringNumberWithoutSpace[0] == ','){
				stringNumberWithoutSpace=strings.Replace(stringNumberWithoutSpace,",","",1)
			}

			stringNumberWithoutSpaceArray := strings.Split(stringNumberWithoutSpace,",")
			for i := 0; i < len(numbers); i++{a,_ := strconv.Atoi(stringNumberWithoutSpaceArray[i]);numbers[i] = a}
			for i := 0; i < len(numbers); i++{
				column := positionsMap[numbers[i]].column
				row := positionsMap[numbers[i]].row
				column = append(column, i+indexOfColumn)
				row = append(row, indexOfRow)
				solution := Position{column, row}
				positionsMap[numbers[i]] = solution
			}
			bingoBoards =append(bingoBoards, numbers)
			indexOfRow++
		}else{
			indexOfColumn+=5
		}
	}

	for _,x := range(selectedNumbers){
		numberX, _ := strconv.Atoi(x)
		if len(positionsMap[numberX].column) > 0 {
			for _,v := range(positionsMap[numberX].column){
				solutionsMapColumn[v]++
				if(solutionsMapColumn[v] == 5){
					lastCalledNumber := numberX
					solution = Solution{index: v, RowOrColumn: "column", lastCalledNumber: lastCalledNumber}
					break;
				}
			}
		}
		if(solution.index > -1){
			break;
		}
		if len(positionsMap[numberX].row) > 0 {
			for _,v := range(positionsMap[numberX].row){
				solutionsMapRow[v]++
				if(solutionsMapRow[v] == 5){
					lastCalledNumber := numberX
					solution= Solution{index: v, RowOrColumn: "row", lastCalledNumber: lastCalledNumber}
					break;
				}
			}
		}
		if(solution.index > -1){
			break;
		}
	}
	var markedNumbers = make(map[int]bool)
	for i := 0; i < len(selectedNumbers); i++{
		n,_ := strconv.Atoi(selectedNumbers[i])
		markedNumbers[n] = true
		if selectedNumbers[i] == strconv.Itoa(solution.lastCalledNumber){
			break;
		}
	}

	sumOfUnmarkedNumbers := 0
	columnToRow := 0;
	if solution.RowOrColumn == "column"{
		columnToRow = 5
	}
	
	var bingoBoard int = solution.index - solution.index%5-columnToRow
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++{
		if !markedNumbers[bingoBoards[bingoBoard][j]]{
			sumOfUnmarkedNumbers += bingoBoards[bingoBoard][j]
		}
		}
		bingoBoard++
	}
	fmt.Println(solution.lastCalledNumber*sumOfUnmarkedNumbers)
	checkError(scanner.Err())

}


func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// idea make a 2d array to save bingo boards, linked list of neighbors 18 <->99<->39
//check for depth of substring as well for