func uniqueOccurrences(arr []int) bool {
    mp := make(map[int]int)
    mpUni := make(map[int]int)
    for _,val := range arr {
        mp[val]++
    }
    for _,val := range mp {
        mpUni[val]++
    }
    for _,val := range mpUni {
        if val > 1 {
            return false
        }
    }
    return true
}