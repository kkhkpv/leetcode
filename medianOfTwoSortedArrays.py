class Solution:
    def findMedianSortedArrays(self, nums1: list[int], nums2: list[int]) -> float:
        min_val = -1000001
        max_val = 1000001
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1

        m = len(nums1)
        n = len(nums2)
        # Binary search
        start = 0
        end = m

        while start <= end:
            nums1_mid = (start + end) // 2
            nums2_mid = (m + n + 1) // 2 - nums1_mid

            if nums1_mid == 0:
                max_of_left_nums1 = min_val
            else:
                max_of_left_nums1 = nums1[nums1_mid - 1]

            if nums1_mid == m:
                min_of_right_nums1 = max_val
            else:
                min_of_right_nums1 = nums1[nums1_mid]

            if nums2_mid == 0:
                max_of_left_nums2 = min_val
            else:
                max_of_left_nums2 = nums2[nums2_mid - 1]

            if nums2_mid == n:
                min_of_right_nums2 = max_val
            else:
                min_of_right_nums2 = nums2[nums2_mid]

            if max_of_left_nums1 <= min_of_right_nums2 and max_of_left_nums2 <= min_of_right_nums1:
                if (m+n) % 2 == 0:
                    return (max(max_of_left_nums1, max_of_left_nums2)+min(min_of_right_nums1, min_of_right_nums2)) / 2
                else:
                    return max(max_of_left_nums1, max_of_left_nums2)
            elif max_of_left_nums1 > min_of_right_nums2:
                end = nums1_mid - 1
            else:
                start = nums1_mid + 1
