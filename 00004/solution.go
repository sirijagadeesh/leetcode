package leetcode

import (
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	allnums := append(nums1, nums2...)
	slices.Sort(allnums)
	cnt := len(allnums)

	if cnt%2 == 0 {
		return float64(allnums[cnt/2]+allnums[(cnt/2)-1]) / 2.0
	}

	return float64(allnums[cnt/2])
}
