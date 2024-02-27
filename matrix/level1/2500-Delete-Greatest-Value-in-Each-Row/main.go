func deleteGreatestValue(grid [][]int) int {
    res := 0
    n := len(grid[0])

    for i := range grid {
        sort.Ints(grid[i])
    }
    for i:= 0;i<n;i++ {
        max := 0
        for j:= 0;j<len(grid);j++ {
            max = getMax(max, grid[j][len(grid[j])-1])
            grid[j] = grid[j][0:len(grid[j])-1] 
        }    
        res += max
    }
    return res
}
func getMax(a int, b int)int{
    if a > b{
        return a
    }
    return b
}