package findMedianSortedArrays

import "fmt"

/**
* leetcode no: 4
*
* @description: 寻找两个有序数组的中位数
*
* @author: toddyhan
*
* @create: 2019/05/17 22:08
**/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		nums1, nums2 = nums2,nums1
		len1, len2 = len2, len1
	}

	iMin := 0
	iMax := len1
	for iMin <= iMax {
		i := (iMin + iMax)/2
		j := (len1 + len2 + 1) / 2 - i
		if i<iMax && nums1[i]<nums2[j-1] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if  j == 0 {
				maxLeft = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					maxLeft = nums1[i-1]
				} else {
					maxLeft = nums2[j-1]
				}
			}
			if (len1+len2) % 2 == 1 {
				return float64(maxLeft)
			}

			var minRight int
			if i == len1 {
				minRight = nums2[j]
			} else if j == len2 {
				minRight = nums1[i]
			} else {
				if nums1[i] > nums2[j] {
					minRight = nums2[j]
				} else {
					minRight = nums1[i]
				}
			}
			return float64(maxLeft + minRight) / 2.0
		}
	}
	return 0.0
}


func main()  {
	nums1 := []int{7,7,7,7,7}
	nums2 := []int{1}
	result := findMedianSortedArrays(nums1,nums2)
	fmt.Println(result)
}