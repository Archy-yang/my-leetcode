package cn

import "fmt"

func findSubstringInWraproundString(p string) int {
	t := map[string]int{}
	m := map[string]int{}

	i, j := 0, 1

	pl := len(p)
	if pl < 1 {
		return 1
	}
	for ;i < pl && j < pl; {
		fmt.Println(p[i])
		fmt.Println(p[j])
		fmt.Println(p[j]-p[i])
		if p[j] - p[j-1] == 1 || p[j-1]-p[j] == 25 {
			j += 1
			continue
		}
		t[string(p[i:j])] = 1
		i = j
		j = j + 1
	}
	if i < pl {
		t[string(p[i:])] = 1
	}

	for s, _ := range t {
		for i:= 0; i < len(s); i++ {
			for j = i; j < len(s); j++ {
				m[s[i:j+1]] ++
			}
		}
	}
	return len(m)
}
