package main

import "fmt"

func twoSum(nums []int, target int) []int {
	count := len(nums)
	for k, _ := range nums {
		for kk := k + 1; kk < count; kk++ {
			if (nums[k] + nums[kk]) == target {
				return []int{k, kk}
			}
		}
	}
	return nil
}

//时间复杂度小于 O(n2)
func twoSum2(nums []int, target int) []int {
	var m = make(map[int]int)
	for k, v := range nums {
		if index, ok := m[target - v]; ok{
			return []int{index, k}
		}else {
			m[v] = k
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum2(nums, target)
	fmt.Println(result)
}
