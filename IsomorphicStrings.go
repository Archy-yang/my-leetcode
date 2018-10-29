package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) < 1 {
		return true
	}
	sM := map[byte]byte{}
	tM := map[byte]byte{}

	sM[s[0]] = t[0]
	tM[t[0]] = s[0]
	for i := 1; i < len(s); i++ {
		if mT, ok := sM[s[i]]; ok && mT != t[i] {
			return false
		}
		if mS, ok := tM[t[i]]; ok && mS != s[i] {
			return false
		}
		sM[s[i]] = t[i]
		tM[t[i]] = s[i]
	}

	return true
}
