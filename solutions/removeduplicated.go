package solutions 

import "sort"

func RemoveDuplicates(nums []int) int {
	duplicated := map[int]int{}
	total := 0
	for i, v := range nums {
		_, ok := duplicated[v]
		if !ok {
			duplicated[v] += 1
			nums[i] = 0
		} else {
			nums[i] = 0
		}
	}
	keys := make([]int, 0, len(duplicated))
	for k := range duplicated {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		nums[total] = k
		total += 1
	}
	return total
}
