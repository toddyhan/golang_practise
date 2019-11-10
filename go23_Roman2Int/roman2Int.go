package Roman2Int

import "strings"

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 11/10/19 7:19 PM
**/

var mapOri = map[string]int {
	"I" : 1,
	"V" : 5,
	"X" : 10,
	"L" : 50,
	"C" : 100,
	"D" : 500,
	"M" : 1000,
}
var mapX = map[string]int {
	"IV" : -2,
	"IX" : -2,
	"XL" : -20,
	"XC" : -20,
	"CD" : -200,
	"CM" : -200,
}

func romanToInt(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		result += mapOri[string(s[i])]
	}
	for k,v := range mapX {
		if strings.Contains(s,k) {
			result += v
		}
	}
	return result
}