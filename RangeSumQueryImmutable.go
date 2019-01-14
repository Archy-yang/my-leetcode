package leetcode

type NumArray struct {
	nums []int
	l int
}


func ConstructorNumArray(nums []int) NumArray {
	return NumArray{
		nums: nums,
		l: len(nums),
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	if this.l < i {
		return 0
	}

	if this.l - 1  < j {
		j = this.l - i
	}
	n := 0
	for k:= i; k <= j; k++{
		n += this.nums[k]
	}

	return n
}


type NumArrayArr struct {
	Arr []int
}


func ConstructorNumArrayArr(nums []int) NumArrayArr {
	sum := 0
	arr := NumArrayArr{}
	for i:=0; i<len(nums); i++{
		sum += nums[i]
		arr.Arr = append(arr.Arr, sum)
	}

	return arr
}


func (this *NumArrayArr) SumRangeArr(i int, j int) int {
	if i == 0 {
		return this.Arr[j]
	}

	return this.Arr[j] - this.Arr[i-1]
}