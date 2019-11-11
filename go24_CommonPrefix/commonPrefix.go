package CommonPrefix

import "strings"

/**
* leetcode no: golang_practise
*
* @description:
*
* @author: toddyhan
*
* @create: 11/10/19 8:20 PM
**/

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _,s := range strs {
		for !strings.HasPrefix(s, prefix) {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}