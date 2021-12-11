package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./input.txt")
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file);
	horizontal := 0
	vertical := 0
	aim := 0
	for(scanner.Scan()){
		instruction := scanner.Text()
		instructionArray := strings.Split(instruction, " ")
		instructionValue,e := strconv.Atoi(instructionArray[1])
		checkError(e)
		switch instructionArray[0]{
		case "forward":
			horizontal += instructionValue
			vertical += aim * instructionValue
		case "down":
			aim += instructionValue
		case "up":
			aim -= instructionValue
		default:
			fmt.Printf("Instruction %s not recognized\n", instructionArray[0])
		}

	}
	checkError(scanner.Err())
	fmt.Println(horizontal*vertical);
}