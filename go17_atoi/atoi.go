package atoi

import (
	"math"
)

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 2019/05/19 19:37
**/

func myAtoi(str string) int {
	var result int
	isBegin := false
	isPositive := true
	isDigitalStart := false
	for index:=0;index< len(str);index++{
		switch str[index] {
		case ' ':
			if isBegin == false {
				continue
			} else {
				return 0
			}
		case '+' :
			if isBegin == false {
				isBegin = true
			} else {
				return 0
			}
			continue
		case '-' :
			if isBegin == false {
				isPositive = false
				isBegin = true

			} else {
				return 0
			}
			continue
		default:
			i := int(str[index] - '0')

			if i < 0 || i > 9 {
				if isBegin == false {
					return 0
				} else {
					if isPositive == true {
						return result
					} else {
						return 0 - result
					}
				}
			} else if i == 0 && isDigitalStart == false {
				if index < len(str)-1 && int(str[index+1]-'0') >= 0 && int(str[index+1]-'0') <= 9 {
					continue
				} else {
					return 0
				}
				//continue
			} else {
				isBegin = true
				isDigitalStart = true
				result = result * 10 + i
				if isPositive == true {
					if result > math.MaxInt32 {
						return math.MaxInt32
					}
				} else {
					if 0-result < math.MinInt32 {
						return math.MinInt32
					}
				}
			}

		}
	}
	if isPositive == true {
		return result
	} else {
		return 0 - result
	}
}