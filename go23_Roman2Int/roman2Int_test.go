package Roman2Int

import (
	"fmt"
	"testing"
)

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 11/10/19 7:33 PM
**/

func Test_rotateWords(t *testing.T) {
	testDatas := []string{
		"I",
		"M",
		"IV",
		"CM",
		"MCMXCIV",
		"",
		"LVIII",
	}
	results := []int{
		1,
		1000,
		4,
		900,
		1994,
		0,
		58,
	}
	for i := 0; i < len(testDatas); i++ {
		result := romanToInt(testDatas[i])
		if result != results[i] {
			fmt.Println("result: ", result)
			t.Fatal("test Wrong!")
		}
	}
}