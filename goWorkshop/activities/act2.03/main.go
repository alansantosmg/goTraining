package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var nums []int

	for i := 0; i < 5; i++ {

		number := rand.Intn(11)

		nums = append(nums, number)

		fmt.Println(nums)

	}

	for i := range nums {

		if i >= 0 {

			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}

			fmt.Println(nums)

		}
		fmt.Println(nums)
	}

}
