package leetcode

import (
	"strings"
	"fmt"
)

var res = [][]string{}
func solveNQueens(n int) [][]string {
	res = [][]string{}
	sh := make([]int, n)
	pi := make(map[int]int)
	na := make(map[int]int)
	dfs(0, n, sh, pi, na, []string{})
	return res
}

//func dfs(row, n int, sh []int, pi, na map[int]int, st []string) {
//	if row >= n {
//		if len(st) == n {
//			res = append(res, st)
//		}
//		return
//	}
//	for k := 0; k < n; k ++ {
//		if sh[k] == 0 && pi[k + row] == 0 && na[row - k] == 0 {
//			sh[k] = 1
//			pi[k + row] = 1
//			na[row - k] = 1
//			t := strings.Repeat(".", k) + "Q" + strings.Repeat(".", n-k -1)
//			stT := append(st, t)
//
//			dfs(row+1, n, sh, pi, na, stT)
//
//			pi[k+row] = 0
//			na[row - k] = 0
//			sh[k] = 0
//		}
//	}
//}

func dfs(row, n int, sh []int, pi, na map[int]int, st []string) {
	if row >= n {
		if len(st) == n {
			res = append(res, st)
			if len(res) == 28 {
				fmt.Println(1)
			}
		}
		return
	}
	for k := 0; k < n; k ++ {
		if sh[k] == 0 && pi[k + row] == 0 && na[row - k] == 0 {
			//t1 := make([]string, len(st))
			//copy(t1, st) // go的slice是指向的一个地址，append的时，如果不重新分配，那么会将原来的值覆盖。
			t1 := st
			sh[k] = 1
			pi[k + row] = 1
			na[row - k] = 1
			t := strings.Repeat(".", k) + "Q" + strings.Repeat(".", n-k -1)
			t1 = append(t1, t)

			dfs(row+1, n, sh, pi, na, t1)

			pi[k+row] = 0
			na[row - k] = 0
			sh[k] = 0
		}
	}
}


func solveNQueens1(n int) [][]string {
	res = [][]string{}
	sh := make([]int, n)
	pi := make([]int, n)
	na := make([]int, n)
	dfs1(1, n, sh, pi, na, []string{})
	return res
}

func dfs1(j, n int, sh, pi, na []int, st []string) {
	if j > n {
		if len(st) == n {
			res = append(res, st)
		}
		return
	}
	str := ""
	for k := 0; k < n; k ++ {
		if sh[k] == 0 && pi[k] == 0 && na[k] == 0 {
			t1 := make([]int, n)
			t2 := make([]int, n)
			t3 := make([]int, n)
			t4 := make([]string, len(st))
			copy(t1, sh)
			copy(t2, pi)
			copy(t3, na)
			copy(t4, st)
			t5 := str
			sh[k] = 1
			pi[k] = 1
			pi = append(pi, 0)
			pi = pi[1:]
			na[k] = 1
			na = append([]int{0}, na[0:n-1]...)
			str += "Q"
			str += strings.Repeat(".", n-k -1)
			st = append(st, str)

			dfs1(j+1, n, sh, pi, na, st)

			copy(sh, t1)
			copy(pi, t2)
			copy(na, t3)
			st = make([]string, len(t4))
			copy(st, t4)
			str = t5
		}
		str += "."
	}
}