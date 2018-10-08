package leetcode

import "testing"

var dataProvider = map[int]string{
	1: "1",
	2: "11",
	3: "21",
	4: "1211",
	5: "111221",
	6: "312211",
	7: "13112221",
	8: "1113213211",
}
func TestCountAndSay(t *testing.T) {
	for num, str := range(dataProvider) {
		re := countAndSay(num)

		if (re != str) {
			t.Errorf("input %v expect value %v actaul %v", num, str, re)
		}
	}
}
