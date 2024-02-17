func sumIndicesWithKSetBits(nums []int, k int) int {
    count := 0

    for i:=0;i<len(nums);i++ {
        if(countBits(i) == k){
            count += nums[i]
        }
    }

    return count
}

func countBits(i int)int{
    count := 0
    for i != 0{
        count += i & 1
        i >>= 1
    }
    return count
}