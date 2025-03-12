package leetcode_ans

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l := len(nums1) + len(nums2)
    res := make([]int, l)
    i1 := 0
    i2 := 0
    for i := 0; i < l; i++ {
        if i1 < len(nums1) && i2 < len(nums2) {
            if nums1[i1] <= nums2[i2] {
                res[i] = nums1[i1]
                i1++
            } else {
                res[i] = nums2[i2]
                i2++
            }
        } else if i1 < len(nums1) {
            res[i] = nums1[i1]
            i1++
        } else {
            res[i] = nums2[i2]
            i2++
        }
    }
    if l % 2 == 0 {
        return float64(res[l/2] + res[l/2-1]) / 2
    } else {
        return float64(res[l/2])
    }
}