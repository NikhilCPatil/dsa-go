func sumCounts(nums []int) int {
    if(len(nums)== 1){
        return 1
    }
    ans := 0
    for i := range nums {
        s := make([]int, 101)
        cnt := 0
        for _,x := range nums[i:] {
            s[x]++
            if s[x] == 1 {
                cnt++
            }
            ans += cnt * cnt
        }   
    }
    return ans
}