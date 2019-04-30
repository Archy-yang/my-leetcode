package leetcode

// 超时。回溯太多
func isMatch(s string, p string) bool {
	sl := len(s)
	pl := len(p)
	if sl == pl && sl == 0 {
		return true
	}
	if pl == 1 && p[0] == '*' {
		return true
	}

	if sl == 0 && pl > 0 && p[0] != '*' {
		return false
	}
	if sl > 0 && pl == 0 {
		return false
	}

	if p[0] == '*' {
		if sl == 0 {
			return isMatch(s, p[1:])
		}
		for i := 0; i < sl; i ++ {
			if isMatch(s[i:], p[1:]) {
				return true
			}
		}
		return false
	} else if p[0] == '?' ||  p[0] == s[0] {
		return isMatch(s[1:], p[1:])
	} else {
		return false
	}
}

// 只对最后一个'*'进行回溯
func isMatch2(s string, p string) bool {
	var star int = -1
	pp, sp, stp := 0, 0, 0
	lp, ls := len(p), len(s)
	for sp < ls {
		if pp < lp {
			// 记录最后一个 * 的位置
			if p[pp] == '*' {
				star = pp
				stp = sp
			}
			if p[pp] == '?' || p[pp] == s[sp] {
				pp++
				sp++
				continue
			}
		}
		// 对 * 进行匹配，从0个匹配开始累加，回溯
		if star >= 0 {
			pp = star + 1
			sp = stp
			stp++
			continue
		}
		return false
	}

	for pp < lp {
		if p[pp] != '*' {
			break
		}
		pp++
	}

	return pp == lp
}

// 动态规划
// DP[i][j] 代表 s的i长度字符与 p 的j长度字符是否匹配
// 对于非 * 字符，那么 DP[i][j] = DP[i-1][j-1]
// 对于 * 字符，会有三种情况
// 1:p 增加了一个字符，s没有增加（* 匹配的是空)
// 2: p没有增加，s增加。 * 匹配了s中多个字符
// 3: p增加了一个，s增加了一个。 * 匹配了s中的一个字符
// 所以： DP[i][j] = DP[i][j-1] || DP[i-1][j] || DP[i-1][j-1]
func isMatch3(s string, p string) bool {
	dp := make([][]bool, len(s) + 1)
	pFirst := make([]bool, len(p)+1)
	pFirst[0] = true
	for i := 0; i < len(p); i++ {
		pFirst[i+1] = pFirst[i] && p[i] == '*'

		if !pFirst[i+1] {
			break
		}
	}
	dp[0] = pFirst
	for i := 0; i < len(s); i++ {
		dp[i + 1] = make([]bool, len(p) + 1)
	}

	for i := 0; i < len(s); i ++ {
		for j:=0; j < len(p); j ++ {
			if p[j] == '?' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				dp[i+1][j+1] = dp[i][j] || dp[i+1][j] || dp[i][j+1]
			}
		}
	}

	return dp[len(s)][len(p)]
}
