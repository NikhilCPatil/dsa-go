func runningSum(nums []int) []int {
    sum := 0
    for i:=0;i<len(nums);i++{
        sum = sum + nums[i]
        nums[i] = sum
    }
    return nums
}