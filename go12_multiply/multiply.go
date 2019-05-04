package main

import (
	"fmt"
)

//leetcode 43

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	if l1 ==0 || l2 == 0 {
		return ""
	}
	if num1=="0" || num2 == "0" {
		return "0"
	}
	result:=make([]byte, l1+l2)
	for i:=l1-1;i>=0; i-- {
		for j:=l2-1;j>=0;j-- {
			no1 := num1[i]-'0'
			no2 := num2[j] - '0'
			mul := no1 * no2
			mul += result[i+j+1]
			result[i+j] += mul/10
			result[i+j+1] = mul%10
		}
	}
	zeroIndex := 0
	for index,val := range result {
		if val != 0 {
			if index != 0 {
				zeroIndex = index - 1
				break
			} else {
				zeroIndex = -1
				break
			}

		}
	}
	result = result[zeroIndex+1:]
	for i:=0;i< len(result);i++ {
		result[i] += '0'
	}
	return string(result)
}


func main()  {
	s1 := "999"
	s2 := "9990"
	res := multiply(s1,s2)
	fmt.Println(res)
}