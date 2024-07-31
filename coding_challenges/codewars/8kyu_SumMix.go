/*
* https://www.codewars.com/kata/57eaeb9578748ff92a000009/train/go
*
* Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
*
* Return your answer as a number.
*
*/

package kata

import (
	"strconv"
)

func SumMix(arr []any) int {
	sum := 0
	for _, x := range arr {
		switch value := x.(type) {
		case string:
			num, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				continue
			}
			sum += int(num)
		case int:
			sum += value
		}
	}
	return sum
}
