package main

import "fmt"

type Solution struct{}

func (s Solution) twoSum(nums []int, target int) []int {
	res := make([]int, 0)
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if target == nums[i]+nums[j] {
				res = append(res, i)
				res = append(res, j)
			}
		}
	}

	return res
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 17
	sol := Solution{}
	res := sol.twoSum(nums, target)
	fmt.Println(res)
}
