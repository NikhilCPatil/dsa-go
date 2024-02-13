func getConcatenation(nums []int) []int {
    
    // var res []int;
    len := len(nums); 
    for i:=0; i< len; i++ {
        nums = append(nums, nums[i])    
    }
    return nums;
}