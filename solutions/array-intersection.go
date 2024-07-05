package solutions

func Intersect(nums1 []int, nums2 []int) []int {
	var result []int
	mapa := make(map[int]int)

	for _, v := range nums1 {
		mapa[v]++

	}
	for _, v := range nums2 {
		if mapa[v] > 0 {
			result = append(result, v)
			mapa[v]--
		}
	}
	return result
}
