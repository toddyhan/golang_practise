package reverseInt

import "math"

/**
* leetcode no: 7
*
* @description: easy,整数反转
*
* @author: toddyhan
*
* @create: 2019/05/19 18:54
**/

func reverse(x int) int {
	var result int
	maxTemp := math.MaxInt32 / 10
	for x != 0 {
		pop := x % 10
		x = x / 10
		if result > maxTemp || (result == maxTemp && pop > 7) {
			return 0
		}
		if result < (0 - maxTemp) || (result == (0-maxTemp) && pop < -8) {
			return 0
		}
		result = result * 10 + pop
	}
	if x < 0 {
		result = 0 - result
	}
	return result
}

