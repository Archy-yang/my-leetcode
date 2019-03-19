package leetcode

func uniquePaths(m int, n int) int {
	var dp [][]int
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}

	dp[0][0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j ++ {
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
		}
	}

	return dp[m-1][n-1];
}
// discuss中，有个O(1)的解法。很巧妙的将问题看成一个排列问题。当m>1,n>1时，例，7，3。就是下移两步，右移6步，一同移动8步，那么就是从8步里取
// 两步为下移。一个排列组合的问题。  8!/(2! * 6!) 厉害~~。
