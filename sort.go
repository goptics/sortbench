package sortbench

func Sort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low, high int) {
	if low < high {
		pi := partition(nums, low, high)
		quickSort(nums, low, pi-1)
		quickSort(nums, pi+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high]
	i := low
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[high] = nums[high], nums[i]
	return i
}
