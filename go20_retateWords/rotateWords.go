package retateWords

/**
* leetcode no:
*
* @description: rotate word for a instence, for example, "I am a student." rotate to "student. a am I"
*
* @author: toddyhan
*
* @create: 7/20/19 10:02 AM
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

func rotateWords(ori string) string {
	if len(ori) == 0 {
		return ""
	}
	spacePos := make([]int, 0)
	oriArray := []byte(ori)
	for index, c := range oriArray {
		if c == ' ' {
			spacePos = append(spacePos, index)
		}
	}
	start := 0

	for i, pos := range spacePos {
		if start == pos {
			start = pos + 1
		} else {
			end := pos
			reverseString(oriArray[start:end])
			start = pos + 1
		}

		if i == len(spacePos)-1 {
			reverseString(oriArray[start:])
		}
	}
	reverseString(oriArray)
	return string(oriArray)
}
