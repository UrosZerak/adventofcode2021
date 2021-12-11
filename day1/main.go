package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func errCheck(e error){
	if e != nil {
		panic(e);
	}
}

func main() {
	var depth []int;
	var increments int = 0;
	file, err := os.Open("./input.txt")
	errCheck(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for(scanner.Scan()){
		fileDepth, err := strconv.Atoi(scanner.Text())
		errCheck(err)
		depth = append(depth, fileDepth)
	}
	lenghtOfDepthArray := len(depth) -1

	errCheck(scanner.Err())

	for i := 0; i < lenghtOfDepthArray - 2; i++ {
		firstThreeSum := depth[i] + depth[i+1] + depth[i+2]
		secondThreeSum := depth[i+1] + depth[i+2] + depth[i+3]
		if(firstThreeSum < secondThreeSum){
			increments++;
		}
	}

	fmt.Println(increments)

}