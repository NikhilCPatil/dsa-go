func findMissingAndRepeatedValues(grid [][]int) []int {
    mp := make(map[int]int)
    missing := 0
    repeted := 0

    i := 1
    for i <= len(grid) * len(grid){
        mp[i] = 0   
        i++ 
    }

    for i:= range grid{
        for _,val := range grid[i]{
            mp[val]++
        }
    }

    for i:=1;i<=len(grid)*len(grid);i++{
        if mp[i] == 2 {
            repeted = i
        }
        if mp[i] == 0{
            missing = i
        }
    }

    return []int{repeted, missing}
}