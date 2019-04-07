package go01_stringreplace

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