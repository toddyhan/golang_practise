package main

//leetcode 15

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && (nums[i] == nums[i-1]) {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		target := 0 - nums[i]
		for {
			if left >= right {
				break
			}
			if nums[left]+nums[right] < target {
				left++
			} else if nums[left]+nums[right] > target {
				right--
			} else {
				resTmp := []int{nums[i], nums[left], nums[right]}
				result = append(result, resTmp)
				left++
				for {
					if (left < right) && (nums[left] == nums[left-1]) {
						left++
					} else {
						break
					}
				}
				right--
				for {
					if (right > left) && (nums[right] == nums[right+1]) {
						right--
					} else {
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	//testData := []int{-1,0,1,2,-1,-4}
	testData2 := []int{-2, 0, 1, 1, 2}
	result := threeSum(testData2)
	fmt.Println(result)
}
