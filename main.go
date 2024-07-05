package main

import (
	"fmt"

	"github.com/angelcoding06/leetcode-solutions/solutions"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(solutions.TwoSum(nums, target))
}
