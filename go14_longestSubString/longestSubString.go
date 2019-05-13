package main

//Leetcode 3
import (
	"fmt"
	"sort"
)

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	locMap := map[uint8]int{}
	result := make([]int, length)
	result[0] = 1
	start := 0
	for i:=1;i< length;i++ {
		locMap[s[i]]=i
		if s[i] == s[i-1] {
			result[i] = 1
			locMap = map[uint8]int{}
		} else {
			if index,ok:= locMap[s[i]]; !ok {
				result[i] = result[i]+1
			} else {
				if index >= start {
					result[i] = 1
					locMap = map[uint8]int{}
				} else {
					result[i] = result[i]+1
				}
			}
		}

	}
	sort.Ints(result)
	return result[len(result)-1]
}

func main()  {
	testData:="abcabcbb"
	res := lengthOfLongestSubstring(testData)
	fmt.Println(res)
}