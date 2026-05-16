package sortbench

func Sort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	aux := make([]int, n)

	// Phase 1: quad swap - sort blocks of 4
	for i := 0; i < n; i += 4 {
		sortBlock4(nums, i, n)
	}

	// Phase 2: merge blocks in powers of 4 (quadsort's parity merge pattern)
	blockSize := 4
	for blockSize < n {
		for i := 0; i < n; i += blockSize * 2 {
			mid := i + blockSize
			end := min(i+blockSize*2, n)
			if mid < end {
				merge(nums, aux, i, mid, end)
			} else {
				copy(aux[i:end], nums[i:end])
			}
		}
		copy(nums, aux)
		blockSize *= 2
	}
}

func sortBlock4(nums []int, start, n int) {
	end := min(start+4, n)
	length := end - start
	if length <= 1 {
		return
	}
	if length == 2 {
		if nums[start] > nums[start+1] {
			nums[start], nums[start+1] = nums[start+1], nums[start]
		}
		return
	}
	if length == 3 {
		if nums[start] > nums[start+1] {
			nums[start], nums[start+1] = nums[start+1], nums[start]
		}
		if nums[start+1] > nums[start+2] {
			nums[start+1], nums[start+2] = nums[start+2], nums[start+1]
		}
		if nums[start] > nums[start+1] {
			nums[start], nums[start+1] = nums[start+1], nums[start]
		}
		return
	}
	// Sorting network for 4 elements
	if nums[start] > nums[start+1] {
		nums[start], nums[start+1] = nums[start+1], nums[start]
	}
	if nums[start+2] > nums[start+3] {
		nums[start+2], nums[start+3] = nums[start+3], nums[start+2]
	}
	if nums[start] > nums[start+2] {
		nums[start], nums[start+2] = nums[start+2], nums[start]
	}
	if nums[start+1] > nums[start+3] {
		nums[start+1], nums[start+3] = nums[start+3], nums[start+1]
	}
	if nums[start+1] > nums[start+2] {
		nums[start+1], nums[start+2] = nums[start+2], nums[start+1]
	}
}

func merge(src, dst []int, start, mid, end int) {
	// Check if already ordered (adaptive optimization)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
