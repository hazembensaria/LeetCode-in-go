package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3,2,4}, 6))
}

func twoSum(nums []int, target int) []int {
	var tmpNums []int
	i := 0
	j := i + 1
	for i < len(nums) {
		for j < len(nums) {
			if nums[i]+nums[j] == target {

				tmpNums = []int{i, j}
			}
			j++
		}
		i++
		j = i + 1
	}
	return tmpNums

}
