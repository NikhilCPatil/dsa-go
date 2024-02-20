func arithmeticTriplets(nums []int, diff int) int {
    mp := map[int]struct{}{}
	triplet := 0

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = struct{}{}
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]-diff]; ok {
			if _, ok := mp[nums[i]+diff]; ok {
				triplet++
			}
		}
	}
	return triplet
}