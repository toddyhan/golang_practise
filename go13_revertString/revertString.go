package main

//leetcode 344
import "fmt"

func reverseString(s []byte) {
	l := len(s)
	if l == 0 {
		return
	}

	i := 0
	j := l - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	s := []byte{'h', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	fmt.Println(string(s))
}
