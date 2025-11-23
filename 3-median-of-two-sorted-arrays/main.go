package main

import "fmt"


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var merged []int
	 l, r := 0, 0
	 for l < len(nums1) && r < len(nums2) {
		if (nums1[l] < nums2[r]){
			merged = append(merged, nums1[l])
			l++
		} else if (nums1[l] >= nums2[r]) {
			merged = append(merged, nums2[r])
			r++
		}
	 }
	 if l < len(nums1){
	 	merged = append(merged, nums1[l:]...)
	 }
	 if r < len(nums2){
	 	merged = append(merged, nums2[r:]...)
	 }
	 new_len := len(merged)
	 if (new_len % 2 == 1){
		return float64(merged[(new_len)/2])
	 } else {
		mid_point := new_len / 2
		return float64(merged[mid_point-1] + merged[mid_point]) / 2
	 }   
}

func main()  {
	var num1 []int = []int{1,3}
	var num2 []int = []int{2}

	fmt.Println(findMedianSortedArrays(num1, num2))
}
