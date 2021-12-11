package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)
	var diagnosticReport []string
	for scanner.Scan(){
		diagnosticReport = append(diagnosticReport, scanner.Text())
	}
	oxygen := depthSearch(diagnosticReport,0, false)
	co := depthSearch(diagnosticReport,0, true)

	oxygenDecimal, _ := strconv.ParseInt(oxygen, 2, 64)
	co2Decimal, _ := strconv.ParseInt(co, 2, 64)
	fmt.Println(oxygenDecimal*co2Decimal)
	
	checkError((scanner.Err()))

}

// recursive operation with first input array of matching strings
func depthSearch(diagnosticReport []string, index int, reverse bool) string {
	if(len(diagnosticReport) == 1){
		return diagnosticReport[0]
	}
	occurence := 0
	for i := 0; i < len(diagnosticReport); i++ {
		if string(diagnosticReport[i][index]) == "1" {
			occurence++
		} else {
			occurence--
		}
	}
 	matchingBit:= "1"
	if !reverse {
	if occurence < 0 {
		matchingBit = "0"
	}
	}else{
		matchingBit = "0";
	if occurence < 0{
			matchingBit = "1"
		}
}
	return depthSearch(returnNewArray(diagnosticReport,index,matchingBit),index+1, reverse);
}

func returnNewArray(diagnosticReport []string, index int, matchingBit string) []string{
	var newDiagnosticReport []string;
	for i := 0; i < len(diagnosticReport); i++{
		if(string(diagnosticReport[i][index]) == matchingBit){
			newDiagnosticReport = append(newDiagnosticReport, diagnosticReport[i])
		}
	}
	return newDiagnosticReport
}


func checkError(e error) {
	if e != nil {
		panic(e)
	}
}