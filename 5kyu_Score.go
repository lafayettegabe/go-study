/*
* https://www.codewars.com/kata/5270d0d18625160ada0000e4/train/go
*
* Greed is a dice game played with five six-sided dice. Your mission, should you choose to accept it, is to score a throw according to these rules. You will always be given an array with five six-sided dice values.
*
*  Three 1's => 1000 points
*  Three 6's =>  600 points
*  Three 5's =>  500 points
*  Three 4's =>  400 points
*  Three 3's =>  300 points
*  Three 2's =>  200 points
*  One   1   =>  100 points
*  One   5   =>   50 point
*
* A single die can only be counted once in each roll. For example, a given "5" can only count as part of a triplet (contributing to the 500 points) or as a single 50 points, but not both in the same roll.
*
* Example scoring
*
*  Throw       Score
*  ---------   ------------------
*  5 1 3 4 1   250:  50 (for the 5) + 2 * 100 (for the 1s)
*  1 1 1 3 1   1100: 1000 (for three 1s) + 100 (for the other 1)
*  2 4 4 5 4   450:  400 (for three 4s) + 50 (for the 5)
*
 */

package kata

func Count(list []int) map[int]int {
	counts := make(map[int]int)
	for _, value := range list {
		counts[value]++
	}
	return counts
}

func Score(dice [5]int) int {
	var counts map[int]int = Count(dice[:])
	total := 0
	
	for key, value := range counts {
		switch key {
		case 1:
			total += (1000 * (value / 3)) + (100 * (value % 3))
		case 2:
			total += (200 * (value / 3))
		case 3:
			total += (300 * (value / 3))
		case 4:
			total += (400 * (value / 3))
		case 5:
			total += (500 * (value / 3)) + (50 * (value % 3))
		case 6:
			total += (600 * (value / 3))
		}
	}
	return total
}
