package retateWords

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
* @create: 7/20/19 10:12 AM
**/

func Test_rotateWords(t *testing.T) {
	testDatas := []string{
		"I am ab student.",
		" student. ",
		"student ",
		" student",
		"   ",
		"",
	}
	results := []string{
		"student. ab am I",
		" student. ",
		" student",
		"student ",
		"   ",
		"",
	}
	for i := 0; i < len(testDatas); i++ {
		result := rotateWords(testDatas[i])
		if result != results[i] {
			fmt.Println("result: ", result)
			t.Fatal("test Wrong!")
		}
	}
}
