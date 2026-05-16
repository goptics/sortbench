package sortbench

func Sort(nums []int) {
	n := len(nums)
	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			key := nums[i]
			j := i
			for j >= gap && nums[j-gap] > key {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = key
		}
		gap /= 2
	}
}
