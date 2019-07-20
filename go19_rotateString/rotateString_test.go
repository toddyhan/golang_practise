package rotateString

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
* @create: 2019/07/20 09:58
**/

func Test_rotateString(t *testing.T) {
	testDatas :=[]string{
		"123456",
		"aabbcc",
		"abcddd",
		"",
		"qwerty",
	}
	testPos :=[]int {
		3,
		6,
		0,
		5,
		5,
	}
	results := []string{
		"561234",
		"abbcca",
		"bcddda",
		"",
		"qwerty",
	}
	for i:=0;i< len(testDatas);i++ {
		result := rotateString(testDatas[i], testPos[i])
		if result != results[i] {
			fmt.Println("result is : ",result)
			t.Fatal("test Wrong")
		}
	}
}