func kidsWithCandies(candies []int, extraCandies int) []bool {
    res :=  make([]bool, len(candies))
    for i := 0;i<len(candies);i++ {
        currentHighest := candies[i] + extraCandies
        res[i] = true      
        for j := 0;j<len(candies);j++{
            if currentHighest < candies[j]{
               res[i] = false
            }
        }
    }
    return res
}