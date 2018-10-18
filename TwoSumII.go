package leetcode

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for {
		if i > j {
			break
		}
		if numbers[i] + numbers[j] == target {
			return []int{i+1, j+1}
		}
		if numbers[i] + numbers[j] > target {
			j--
			continue
		}
		if numbers[i] + numbers[j] < target {
			i++
			continue
		}
	}
	return []int{}
}
