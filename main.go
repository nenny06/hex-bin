package main

import (
	"fmt"
)

//	func binToDec(input string, base int) int {
//		decimal := 0
//		power := 0
//		var value int
//		for char := len(input) - 1; char >= 0; char-- {
//			if char >= '0' || char <= '9' {
//				value = int(char - '0')
//			} else if char >= 'A' || char <= 'F' {
//				value = int(char-'A') + 10
//			} else if char >= 'a' || char <= 'f' {
//				value = int(char-'a') + 10
//			}
//			decimal += value * int(math.Pow(float64(base), float64(power)))
//			power++
//		}
//		return decimal
//	}// modified code
func binToDec(input string, base int) int {

	if base != 16 && base != 2 { // handles base input error

		fmt.Println("")
		return 0
	}

	decimal := 0 // where to store num converted to base

	for i := 0; i < len(input); i++ { // loop through input
		ch := input[i] // initializing and assigning ch to input inde to help with comparision
		var value int

		if ch >= '0' && ch <= '9' {
			value = int(ch - '0')
		} else if ch >= 'A' && ch <= 'F' {
			value = int(ch-'A') + 10
		} else if ch >= 'a' && ch <= 'f' {
			value = int(ch-'a') + 10
		}

		decimal = decimal*base + value
	}

	return decimal
}

func main() {
	fmt.Println(binToDec("11A", 16))
	fmt.Println("")
	fmt.Println(binToDec("101", 2))
}
