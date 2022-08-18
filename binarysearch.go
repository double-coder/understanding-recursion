package main

func binarysearch(nums []int, target int, left, right int) bool {
	if left > right {
		return false
	}
	mid := left + (right-left)/2

	if nums[mid] == target {
		return true
	}

	if nums[mid] < target {
		return binarysearch(nums, target, left, mid-1)
	}

	return binarysearch(nums, target, mid+1, right)
}
