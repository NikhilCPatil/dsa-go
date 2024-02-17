func decompressRLElist(nums []int) []int {
    deComp := []int{}
    for i:=0;i<len(nums);i+=2{
       freq := nums[i]
       val := nums[i+1]
       for j:=0;j<freq;j++{
           deComp = append(deComp, val)
       }
    }
    return deComp
}