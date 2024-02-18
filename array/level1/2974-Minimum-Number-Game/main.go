func numberGame(nums []int) []int {
    arr := make([]int,0,len(nums))
    sort.Ints(nums)
    for i:=0;i<len(nums);i +=2{
        arr = append(arr,nums[i+1],nums[i])
    } 
    return arr
}