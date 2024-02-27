func countNegatives(grid [][]int) int {
    count := 0
    for i:= range grid{
        for j:=len(grid[i])-1;j>=0;j--{
            if grid[i][j] < 0{
                count++
            }
        }
    }
    return count
}