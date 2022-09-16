package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// 1358 9549 9391 4435
// Receive card detailas input
func enteredNumber() string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the card number you need to validate: ")
	input, _ := reader.ReadString('\n')
	card := strings.ReplaceAll(input, " ", "")
	return card
	
}

// Reverse string to allow for odd & even card number to be process 
func reverseString(str string) string{
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
	   byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	
	Data := string(byte_str)
	return Data
	
 }


// Convert string to integer
func convertToInt(str string) int{
	newInt, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Attention! ", err)
	} 
	return newInt
}

// sum values
func sumValues(num int, total int) {
	total += num
}

// Seperate between out every 2nd digit and the remaining
func seperateNumber() {

}

// Check if card is valid
func validateCard(odds int, evens int) {
	validate := odds + evens
	if validate %10 == 0 {
		fmt.Println("CARD NUMBER PROVIDED IS VALID")
	} else {
		fmt.Println("NOT A VALID CARD NUMBER")
	}

}

func main() {
	
	
	var strOfNums string = reverseString(enteredNumber())
	



	sumOfEvens := 0
	sumOfOdds := 0
	doubled := 0
	calculated := 0


	for i, v := range strOfNums {
		
		// Skip index 0 as it contains \n and causes an error
		if i == 0 {
			continue;
		}
		if i % 2 == 0 {
			doubled = convertToInt(string(v)) * 2
			if doubled > 9 {
				calculated = doubled - 9
			}else{
				calculated = doubled
			}
			sumOfOdds += calculated
		}
		if i % 2 == 1 {
			sumOfEvens += convertToInt(string(v))
		}
		
	}



	
	validateCard(sumOfOdds, sumOfEvens)
	
}