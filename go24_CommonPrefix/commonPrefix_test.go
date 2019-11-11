package CommonPrefix

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
* @create: 11/10/19 8:26 PM
**/

func Test_longestCommonPrefix(t *testing.T) {
	testDatas := [][]string{
		{"flower", "flow", "flay"},
		{"a", "b", "c"},
		{"c","acc","ccc"},
	}
	results := []string{
		"fl",
		"",
		"",
	}
	for i := 0; i < len(testDatas); i++ {
		result := longestCommonPrefix(testDatas[i])
		if result != results[i] {
			fmt.Println("result: ", result)
			t.Fatal("test Wrong!")
		}
	}
}