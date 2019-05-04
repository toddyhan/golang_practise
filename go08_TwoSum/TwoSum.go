package main

//leetcode 1
import "fmt"

//利用map，空间换时间，O(n)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var result = map[int]int{}
	for index, value := range nums {
		if _, ok := result[target-value]; !ok {
			result[value] = index
		} else {
			return []int{result[target-value], index}
		}
	}
	return []int{}
}

//另一解法，把target-i的index存入，可以保证一旦找到立即返回
func twoSum2(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	var result = map[int]int{}
	for index, value := range nums {
		if index2, ok := result[value]; ok {
			return []int{index2, index}
		}
		result[target-value] = index
	}
	return []int{}
}
func main() {
	var testData = []int{3, 2, 4, 2, 3, 4}
	result := twoSum(testData, 6)
	fmt.Println(result)
}
