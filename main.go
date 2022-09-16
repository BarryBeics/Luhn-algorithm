package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


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
	fmt.Println(Data)
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
	return total
}

// wfe
func calculate(baz string) {
	evens := 0
	odds := 0
	doubled := 0
	calculated := 0
	for i, v := range baz {
		
		// Skip index 0 as it contains \n and causes an error
		if i == 0 {
			continue;
		}
		if i % 2 == 0 {
			odds = convertToInt(string(v))
			doubled = odds * 2
			if doubled > 9 {
				calculated = doubled - 9
			}else{
				calculated = doubled
			}

			//odds += convertToInt(string(v))
			fmt.Println("2nd's", odds)
		}
		if i % 2 == 1 {
			
			evens = convertToInt(string(v))
			evens = evens * 2
			fmt.Println(evens)
		}
		
	}
	calculate(strOfNums)
	fmt.Println("total of the odds: ", odds)
	fmt.Println(calculated)
}

func main() {
	
	
	var strOfNums string = reverseString(enteredNumber())
	
	
	
}