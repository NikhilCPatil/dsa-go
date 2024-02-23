func findDifference(nums1 []int, nums2 []int) [][]int {
    res := make([][]int, 2)
    mp1 := make(map[int]int)
    mp2 := make(map[int]int)

    for _,num := range nums1{
        mp1[num]++
    }

    for _,num := range nums2{
        mp2[num]++
    }

    for i,_ := range mp1{
       if mp2[i] == 0 {
           res[0] = append(res[0], i)
       }
    } 

     for i,_ := range mp2{
       if mp1[i] == 0 {
           res[1] = append(res[1], i)
       }
    } 

    return res

    
}