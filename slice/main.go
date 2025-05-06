package main

import "fmt"

func main() {
	nums := make([]int, 0, 5)
	fmt.Println(nums)
	num2 := make([]int, len(nums))
	copy(num2, nums)
	fmt.Println(cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	fmt.Println(cap(nums))
	fmt.Println(num2)

}
