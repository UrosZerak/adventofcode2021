package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./input.txt")
	checkErr(err)
	powerConsumption := string(data)
	columnConsumption := ""
	cut := 0
	mapOfConsumption := make(map[int]string)


	for i := 0; i < len(powerConsumption); i++{
		if(string(powerConsumption[i]) == "\n"){
			cut = 0;
		}else{
			mapOfConsumption[cut] += string(powerConsumption[i])
			cut++
		}	
	}

	for i := 0; i < len(mapOfConsumption); i++{
		if(countOccurences(mapOfConsumption[i]) > (len(mapOfConsumption[i])-1)/2){
			columnConsumption+="1"
		}else{
			columnConsumption+="0"
		}
	}
	gammaRateDecimal,_ := strconv.ParseInt(columnConsumption, 2, 64)


	for i := 0; i < len(columnConsumption); i++{
		if(string(columnConsumption[i]) == "1"){
			columnConsumption = replaceAtIndex(columnConsumption,'0',i)
		}else{
			columnConsumption = replaceAtIndex(columnConsumption,'1',i)
		}
	}

	epsilonRateDecimal,_ := strconv.ParseInt(columnConsumption, 2, 64)

	fmt.Println(gammaRateDecimal*epsilonRateDecimal)
}

func countOccurences(s string) int{
	return strings.Count(s,"1")
}

func replaceAtIndex(str string, replacement rune, index int) string {
    return str[:index] + string(replacement) + str[index+1:]
}