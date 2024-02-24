func sumOfUnique(nums []int) int {
    mp := make(map[int]int)
    sum := 0
    for _,val := range nums {
        mp[val]++
        if mp[val] == 1 {
            sum += val     
        }
        if mp[val] == 2 {
            sum -=val
        }
    }

    return sum
}