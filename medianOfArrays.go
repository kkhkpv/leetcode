import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	const (
		maxVal int = 1000001
		minVal int = -1000001
	)

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := len(nums1)
	n := len(nums2)
	// Bin search
	start, end := 0, m
	maxOfLeftNums1, minOfRightNums1, maxOfLeftNums2, minOfRightNums2 := 0, 0, 0, 0

	for start < end {
		nums1Mid := (start + end) / 2
		nums2Mid := (m+n+1)/2 - nums1Mid

		if nums1Mid == 0 {
			maxOfLeftNums1 = minVal
		} else {
			maxOfLeftNums1 = nums1[nums1Mid-1]
		}
		if nums1Mid == m {
			minOfRightNums1 = maxVal
		} else {
			minOfRightNums1 = nums1[nums1Mid]
		}
		if nums2Mid == 0 {
			maxOfLeftNums2 = minVal
		} else {
			maxOfLeftNums2 = nums2[nums2Mid-1]
		}
		if nums2Mid == n {
			minOfRightNums2 = maxVal
		} else {
			minOfRightNums2 = nums2[nums2Mid]
		}

		if maxOfLeftNums1 <= minOfRightNums2 && maxOfLeftNums2 <= minOfRightNums1 {
			if (m+n)%2 == 0 {
				return (math.Max(float64(maxOfLeftNums1), float64(maxOfLeftNums2)) + math.Min(float64(minOfRightNums1), float64(minOfRightNums2))) / 2
			} else {
				return math.Max(float64(maxOfLeftNums1), float64(maxOfLeftNums2))
			}
		} else if maxOfLeftNums1 > minOfRightNums2 {
			end = nums1Mid - 1
		} else {
			start = nums1Mid + 1
		}

	}
	return -1
}
