package main
//leetcode 260
import "fmt"

func findSingleInt3 (nums []int) []int {
	var xor int
	for _,i := range nums {
		xor = xor ^ i
	}
	var c uint
	for {

		if xor>>c & 1 == 1 {

			break
		}
		c = c +1
	}
	var a,b int
	for _,i := range nums {
		if i>>c & 1 == 1 {
			a = a ^ i
		} else {
			b = b ^ i
		}
	}
	return []int{a,b}
}


func main() {
	testData1 := []int{1,1,2,2,3,6,3,5,5,4}
	res := findSingleInt3(testData1)
	fmt.Println(res)
}
