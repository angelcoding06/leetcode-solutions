package main

import (
	"fmt"

	"github.com/angelcoding06/leetcode-solutions/solutions"
)

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(solutions.Intersect(nums1, nums2)) // [2, 2]

	nums1 = []int{4, 9, 5}
	nums2 = []int{9, 4, 9, 8, 4}
	fmt.Println(solutions.Intersect(nums1, nums2)) // [4, 9]
}
