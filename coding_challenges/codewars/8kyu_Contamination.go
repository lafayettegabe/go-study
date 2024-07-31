/*
* https://www.codewars.com/kata/596fba44963025c878000039/train/go
*
* An AI has infected a text with a character!!
*
* This text is now fully mutated to this character.
*
* If the text or the character are empty, return an empty string.
* There will never be a case when both are empty as nothing is going on!!
*
* Note: The character is a string of length 1 or an empty string.
*
* Example
* text before = "abc"
* character   = "z"
* text after  = "zzz"
*
*/

package kata

func Contamination(text, char string) string {
	if len(text) == 0 || len(char) != 1 {
		return ""
	}

	response := make([]byte, len(text))
	for i := range response {
		response[i] = char[0]
	}
	return string(response)
}
