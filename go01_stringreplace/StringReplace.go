package go01_stringreplace
/*
题目：
将一个字符串中的空格替换成 "%20"。
Input:
"A B"
Output:
"A%20B"

思路：创建一个比原数组大（数组个数*2）的数组。然后创建两个指针，p1指向原数组的末尾，p2指向新数组的末尾。
然后原数组使用p1向前遍历，如果遇到空格，则新数组从p2处依次向前修改内容为02%；如果遇到的不是空格，则直接将p1处内容放到p2处。
*/
func StringReplace(input string) (output string ,err error) {
	if len(input) == 0 {
		return "", nil
	}
	origin := []byte{}
	origin = []byte(input)
	p1 := len(input)-1
	spaceNum := 0
	for _,i := range origin {
		if i == ' ' {
			spaceNum = spaceNum + 1
		}
	}
	result := make([]byte, len(input)+spaceNum*2)
	p2 := len(result) - 1

	for {
		if p1 < 0 {
			break
		}
		if origin[p1] == ' ' {
			result[p2] = '0'
			p2--
			result[p2] = '2'
			p2--
			result[p2] = '%'
			p2--
		} else {
			result[p2] = origin[p1]
			p2--
		}
		p1--
	}

	return string(result),nil
}
