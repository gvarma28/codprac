package algorithms

func MergeSort(nums []int, start, end int) []int {
	// n -> total range
	n := end - start
	if n <= 1 {
		return nums
	}

	mid := (start + end) / 2
	MergeSort(nums, start, mid)
	MergeSort(nums, mid, end)

	temp := make([]int, n)
	p, q, i := start, mid, 0

	for p < mid && q < end {
		if nums[p] <= nums[q] {
			temp[i] = nums[p]
			p++
		} else {
			temp[i] = nums[q]
			q++
		}
		i++
	}

	for p < mid {
		temp[i] = nums[p]
		p++
		i++
	}
	for q < end {
		temp[i] = nums[q]
		q++
		i++
	}

	copy(nums[start:end], temp)
	return nums
}
