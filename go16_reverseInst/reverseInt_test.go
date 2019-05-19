package reverseInt

import (
	"testing"
)

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 2019/05/19 19:11
**/

func Test_reverseInt(t *testing.T) {
	testData := []int{123, -123}
	resultData := []int{321, -321}
	for i := 0; i< len(testData); i++ {
		result := reverse(testData[i])
		if result != resultData[i] {
			t.Fatal("test wrong")
		}
	}
}
