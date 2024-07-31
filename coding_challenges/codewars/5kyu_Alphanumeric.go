/*
* https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go
*
* In this example you have to validate if a user input string is alphanumeric. The given string is not nil/null/NULL/None, so you don't have to check that.
*
* The string has the following conditions to be alphanumeric:
*
* At least one character ("" is not valid)
* Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
* No whitespaces / underscore
*
 */

package kata

func checkAlpha(letter rune) bool {
	return (letter >= '0' && letter <= '9') || (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z')
}

func alphanumeric(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, letter := range str {
		if !checkAlpha(letter) {
			return false
		}
	}
	return true
}
