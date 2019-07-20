package atoi

import (
	"fmt"
	"math"
	"testing"
)

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 2019/05/19 19:58
**/

func Test_atoi(t *testing.T) {
	testDatas :=[]string{
		"12345",
		"   -54321",
		"word 12345",
		"12345 word",
		"hello",
		"-91283472332",
		"11111111111111111111111111111",
		"3.1415926",
		"+-2",
		"-+1",
		"-0012",
		"+0012",
		"-0012abc",
		"   +0 123",
		"  0000000000012345678",
		"  00000000000",
		"- 234",
	}
	results := []int{
		12345,
		-54321,
		0,
		12345,
		0,
		math.MinInt32,
		math.MaxInt32,
		3,
		0,
		0,
		-12,
		12,
		-12,
		0,
		12345678,
		0,
		0,
	}
	for i:=0;i< len(testDatas);i++ {
		result := myAtoi(testDatas[i])
		fmt.Printf("s is : %s, result is : %d\n", testDatas[i],result)
		if result != results[i] {
			t.Fatal("test Wrong")
		}
	}
}