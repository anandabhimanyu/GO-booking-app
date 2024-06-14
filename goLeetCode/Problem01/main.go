package main

import "fmt"

func main() {
	fmt.Println("Two Sum")
	result := twoSum([]int{3, 5, 6, 13}, 9)
	fmt.Println("result:", result)
}

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		if index, ok := indexMap[difference]; ok {
			return []int{i, index}
		}
		indexMap[nums[i]] = i

	}
	return []int{}

}
