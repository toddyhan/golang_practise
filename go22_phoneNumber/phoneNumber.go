package phoneNumber

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 11/10/19 5:54 PM
**/

func letterCombinations(digits string) []string {
	var phone = map[byte][]byte {
		'2': {'a','b','c'},
		'3': {'d','e','f'},
		'4': {'g','h','i'},
		'5': {'j','k','l'},
		'6': {'m','n','o'},
		'7': {'p','q','r','s'},
		'8': {'t','u','v'},
		'9': {'w','x','y','z'},
	}

}