func smallerNumbersThanCurrent(nums []int) []int {
    res := make([]int, len(nums))
    for i:= 0;i<len(nums);i++{
        curCount := 0
        for j:= 0;j<len(nums);j++{
            if i == j {
                continue
            }
            if (nums[i] > nums[j]){
                curCount++
            }
        } 
        res[i] = curCount   
    }
    return res
}