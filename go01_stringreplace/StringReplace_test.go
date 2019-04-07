package go01_stringreplace

import "testing"

type stringReplaceTest struct {
	in string
	out string
}
var stringReplaceTestData = []stringReplaceTest {
	{" hello world ", "%20hello%20world%20"},
	{"   ", "%20%20%20"},
	{"", ""},
}
func TestStringReplace(t *testing.T) {
	for i, test := range stringReplaceTestData {
		out,err := StringReplace(test.in)
		if nil != err {
			t.Errorf("#%d : error %v", i, err)
			continue
		}
		if out != test.out {
			t.Errorf("#%d : out is: %v, want: %v", i, out, test.out)
		}
	}
}
