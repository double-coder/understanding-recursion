package main

func mergesort(nums []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		mergesort(nums, start, mid)
		mergesort(nums, mid+1, end)
		merge(nums, start, mid, end)
	}
}

func merge(nums []int, start, mid, end int) {
	// build tmp arr to avoid
	temp := make([]int, end-start+1)

	i, j, k := start, mid+1, 0

	// while both sub-array has values
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			k++
			i++
		} else {
			temp[k] = nums[j]
			j++
			k++
		}
	}
	for i <= mid {
		temp[k] = nums[i]
		k++
		i++
	}

	for j <= end {
		temp[k] = nums[j]
		k++
		j++
	}
	for i = start; i <= end; i++ {
		nums[i] = temp[i-start]
	}
}
