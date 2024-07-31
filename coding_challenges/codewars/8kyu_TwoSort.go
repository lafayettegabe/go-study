/*
* https://www.codewars.com/kata/57cfdf34902f6ba3d300001e/train/go
*
* You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.
*
* The returned value must be a string, and have "***" between each of its letters.
*
* You should not remove or add elements from/to the array.
*
 */

package kata

func BubbleSort(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func StarWord(word string) string {
	stared := ""
	for index, char := range word {
		stared += string(char)
		if index+1 != len(word) {
			stared += "***"
		}
	}
	return stared
}

func TwoSort(arr []string) string {
	BubbleSort(arr)
	return StarWord(arr[0])
}
