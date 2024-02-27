func flipAndInvertImage(image [][]int) [][]int {
    for i:= range image{
        j := 0
        k := len(image)-1
        for j <= k {
            image[i][j], image[i][k] = invert(image[i][k]), invert(image[i][j])
            j++
            k-- 
        } 
    }
    return image
}

func invert(a int) int{
    if a == 1 {
        return 0
    }
    return 1
}