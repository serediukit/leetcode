package leetcode_ans

func countGood(nums []int, k int) int64 {
    mp := make(map[int]int)
    left, right := 0, 0
    count := 0
    var ans int64 = 0
	for right < len(nums) {
        count += mp[nums[right]]
        mp[nums[right]]++
        for left < len(nums) && (count - (mp[nums[left]] - 1)) >= k {
            mp[nums[left]]--
            count -= mp[nums[left]]
            left++
        }
        if count >= k {
            ans += int64(left + 1)
        }
        right++
    }
    return ans
}