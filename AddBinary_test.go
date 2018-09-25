package leetcode

import (
	"testing"
	"fmt"
)

var addBinaryTest = []map[string]string{
	{
		"a": "0",
		"b": "0",
		"r": "0",
	},
	{
		"a": "1",
		"b": "0",
		"r": "1",
	},
	{
		"a": "0",
		"b": "1",
		"r": "1",
	},
	{
		"a": "1",
		"b": "1",
		"r": "10",
	},
	{
		"a": "10",
		"b": "1",
		"r": "11",
	},
	{
		"a": "11",
		"b": "1",
		"r": "100",
	},
	{
		"a": "1010",
		"b": "1011",
		"r": "10101",
	},
}

func TestAddBinary(t *testing.T) {
	for _, oneTest := range addBinaryTest{
		act := addBinary(oneTest["a"], oneTest["b"])

		if (len(act) != len(oneTest["r"])) {
			fmt.Printf("error : %v, act %v\n", oneTest, act)
		} else {
			fmt.Printf("success : %v\n", oneTest)
		}
	}

}
