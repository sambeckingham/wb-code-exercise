package main

import (
	"fmt"
)

func main() {
	wordToNumberMap := make(map[int]string)
	wordToNumberMap[0] = "zero"
	wordToNumberMap[1] = "one"
	wordToNumberMap[3] = "three"
	wordToNumberMap[10] = "ten"
	wordToNumberMap[17] = "seventeen"
	wordToNumberMap[20] = "twenty"
	wordToNumberMap[100] = "hundred"
	wordToNumberMap[1000] = "thousand"

	number := 117
	// numberString := strconv.Itoa(number)
	//
	// var numberToTextResult = "test"
	// for _, digit := range numberString {
	//
	// 	// numberToTextResult = fmt.Sprintf("%s%s", numberToTextResult, wordToNumberMap[string(digit)])
	// }

	// fmt.Println(numberToTextResult)

	hundreds := number / 100
	tens := (number - (hundreds * 100) )/ 10
	singles := number - (hundreds * 100) - (tens * 10)

	fmt.Printf("%d hundred\n", hundreds)
	fmt.Printf("%d tens\n", tens)
	fmt.Printf("%d singles\n", singles)

	tensAndSingles := ""
	if (tens >= 10 || tens < 20) {
		tensAndSingles = wordToNumberMap[((tens*10) + singles)]
	}

	numberAsString := ""

	if (hundreds == 0) {
		numberAsString = fmt.Sprintf("%s %s",wordToNumberMap[tens*10], wordToNumberMap[singles])
	} else if (tens == 0 ) {
		numberAsString = fmt.Sprintf("%s hundred and %s", wordToNumberMap[hundreds], wordToNumberMap[singles])
	} else {
		numberAsString = fmt.Sprintf("%s hundred and %s", wordToNumberMap[hundreds], tensAndSingles)
	}

	fmt.Printf(numberAsString)
}