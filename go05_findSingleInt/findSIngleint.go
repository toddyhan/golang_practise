package main
//leetcode 136
import "fmt"

func findSingleInt (nums []int) (int) {
	var result int
	for _,i := range nums {
		result = result ^ i
	}
	return result
}

func main() {
	testData1 := []int{1,2,1,2,3}
	res := findSingleInt(testData1)
	fmt.Println(res)
}
