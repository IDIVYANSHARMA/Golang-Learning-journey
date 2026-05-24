package main

import (
	"fmt"
	"slices"
)

func main() {
	var nums1 = []int {1, 2, 3}
	var nums2 = []int {1, 2, 3}
	fmt.Println(slices.Equal(nums1, nums2))
	
}