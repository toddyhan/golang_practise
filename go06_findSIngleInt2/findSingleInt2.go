package main

import "fmt"

//leetcode 137


func findSingleInt2(nums []int) (int) {
	var a int
	var b int
	for _,i := range nums {
		b = (b ^ i) & (^a)
		a = (a ^ i) & (^b)
	}
	return b

}
func main () {
	testData1 := []int {1,2,1,1,2,3,2}
	res := findSingleInt2(testData1)
	fmt.Println(res)
}