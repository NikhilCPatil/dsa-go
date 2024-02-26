func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for i:=0;i<len(accounts);i++{
        currentTotal := 0
        for j:=0;j<len(accounts[i]);j++{
          currentTotal += accounts[i][j]
        }
        if currentTotal > maxWealth{
            maxWealth = currentTotal    
        }
    }
    return maxWealth
}