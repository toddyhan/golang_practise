package findMedianSortedArrays

import "testing"


/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 2019/05/18 16:38
**/

func Test_findMedianSortedArrays(t *testing.T) {
	testA1 := []int{1,2,3}
	testB1 := []int{4,5,6}
	testA2 := []int{1,3}
	testB2 := []int{2,4,5}
	testA3 := []int{}
	testB3 := []int{1,2,3}

	var result float64

	result = findMedianSortedArrays(testA1,testB1)
	if result != 3.5 {
		t.Errorf("test is wrong, testA is : %v, testB is : %v, result is %v", testA1, testB1, result)
		t.Fatal("test is wrong")
	}

	result = findMedianSortedArrays(testA2,testB2)
	if result != 3.0 {
		t.Errorf("test is wrong, testA is : %v, testB is : %v, result is %v", testA1, testB1, result)
		t.Fatal("test is wrong")
	}

	result = findMedianSortedArrays(testA3,testB3)
	if result != 2.0 {
		t.Errorf("test is wrong, testA is : %v, testB is : %v, result is %v", testA1, testB1, result)
		t.Fatal("test is wrong")
	}
}