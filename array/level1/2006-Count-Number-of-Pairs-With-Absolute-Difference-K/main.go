func countKDifference(nums []int, k int) int {
    pairsCount := 0
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            a := nums[i]
            b := nums[j]
            res := 0
            if(a > b){
                res = a-b
            }else{
                res = b-a
            }

            if res == k {
                pairsCount++
            }
        }   
    }
    return pairsCount
}