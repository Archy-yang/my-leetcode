package leetcode

func missingNumber(nums []int) int {
	largest := 0
	list := make([]bool, len(nums) + 1)

	for _, num := range nums {
		if largest < num {
			largest = num
		}
		list[num] = true
	}

	if largest < len(nums) {
		return len(nums)
	}

	num := 0
	for i := 0; i <= largest; i++ {
		if list[i] == false {
			num = i
			break
		}
	}

	return  num
}
