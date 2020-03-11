package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	nums1, nums2 = nums2, nums1
	fmt.Println(nums1)
	fmt.Println(nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n { // to ensure m<=n
		nums1, nums2 = nums2, nums1
		m,n = n,m
	}
	iMin := 0
	iMax := m
	halfLen := int(math.Floor(float64(m + n + 1) / 2.0))
	for iMin <= iMax {
		i := int(math.Floor(float64(iMin + iMax) / 2.0))
		j := halfLen - i
		fmt.Println(i,j)
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1 // i is too small
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1 // i is too big
		} else { // i is perfect
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0

			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = Min(nums2[j], nums1[i])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}

	return 0.0
}

//Max lalala
func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

//Min lalala
func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
