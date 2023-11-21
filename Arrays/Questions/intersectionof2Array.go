package questions

// Given two integer arrays nums1 and nums2, return an
// array of their intersection. Each element in the result must be
// unique and you may return the result in any order.
func Intersection(nums1 []int, nums2 []int) []int {

	numSet := make(map[int]bool)
	for _, num := range nums1 {
		numSet[num] = true
	}

	var result []int

	for _, num := range nums2 {
		if numSet[num] {
			result = append(result, num)

			delete(numSet, num)
		}
	}

	return result
}
