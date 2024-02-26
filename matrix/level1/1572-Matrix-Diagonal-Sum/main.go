func diagonalSum(mat [][]int) int {
    ret := 0
    n := len(mat)
    for i:=0;i<len(mat);i++{
        ret += mat[i][i] + mat[i][n-i-1] 
    }

    if n % 2 == 1  {
            ret -= mat[n/2][n/2]
        }  

    return ret
}