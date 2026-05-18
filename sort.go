package sortbench

func Sort(nums []int) {
	if len(nums) < 2 {
		return
	}
	aux := make([]int, len(nums))
	mergeSort(nums, aux, 0, len(nums))
}

func mergeSort(nums, aux []int, start, end int) {
	if end-start < 2 {
		return
	}
	if end-start == 2 {
		if nums[start] > nums[start+1] {
			nums[start], nums[start+1] = nums[start+1], nums[start]
		}
		return
	}

	mid := start + (end-start)/2
	mergeSort(nums, aux, start, mid)
	mergeSort(nums, aux, mid, end)
	merge(nums, aux, start, mid, end)
}

func merge(src, dst []int, start, mid, end int) {
	if src[mid-1] <= src[mid] {
		copy(dst[start:end], src[start:end])
		return
	}

	i, j, k := start, mid, start
	for i < mid && j < end {
		if src[i] <= src[j] {
			dst[k] = src[i]
			i++
		} else {
			dst[k] = src[j]
			j++
		}
		k++
	}
	for i < mid {
		dst[k] = src[i]
		i++
		k++
	}
	for j < end {
		dst[k] = src[j]
		j++
		k++
	}
}
