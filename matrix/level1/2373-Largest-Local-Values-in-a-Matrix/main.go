func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    maxLocal := make([][]int, n-2)

    for i := range maxLocal {
        maxLocal[i] = make([]int, n-2)
    }

    for i := 0;i<n-2;i++{
        for j:=0;j<n-2;j++{
            //find highest for 3*3
            highest := 0
            for x:= i;x<i+3;x++{
                for y:=j;y<j+3;y++{
                    if grid[x][y] > highest{
                        highest = grid[x][y]
                    }
                }
            }
            maxLocal[i][j] = highest
        }
    }

    return maxLocal
}
