package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	if length == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	closest := math.MaxInt32
	for i := 0; i < length-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := length - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			abs := int(math.Abs(float64(sum - target)))
			absClose := int(math.Abs(float64(closest - target)))
			if abs < absClose {
				closest = sum
			}
			if sum > target {
				right--
				for left < right {
					if nums[right] == nums[right+1] {
						right--
					} else {
						break
					}
				}
			} else {
				left++
				for left < right {
					if nums[left] == nums[left-1] {
						left++
					} else {
						break
					}
				}
			}
		}
	}
	return closest
}

func main() {
	var testData = []int{-1, 2, 1, -4}
	result := threeSumClosest(testData, 1)
	fmt.Println(result)
}
