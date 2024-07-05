package solutions

func TwoSum(nums []int, target int) []int {
	var result = make([]int, 2)

	mapa := make(map[int]int)
	for i, v := range nums {
		diferencia := target - v
		if v, ok := mapa[diferencia]; ok {
			result[0] = v
			result[1] = i

		}
		mapa[v] = i
	}

	return result

}
