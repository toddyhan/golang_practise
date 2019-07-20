package rotateString

/**
* leetcode no: 编程之法-字符串的旋转
*
* @description: abcdef旋转为defabc
*
* @author: toddyhan
*
* @create: 2019/07/20 09:37
**/

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

func rotateString(ori string, pos int) string {
	if len(ori) == 0 {
		return ""
	}
	rpos := pos % len(ori)
	oriArray := []byte(ori)
	reverseString(oriArray[0:rpos+1])
	reverseString(oriArray[rpos+1:])
	reverseString(oriArray)
	return string(oriArray)
}