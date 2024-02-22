func findIntersectionValues(nums1 []int, nums2 []int) []int {
    res := []int{}
    mp1 := make([]bool, 101)
    mp2 := make([]bool, 101)

    for _,num := range nums1{
        mp1[num] = true
    }

    for _,num := range nums2{
        mp2[num] = true
    }

    ans1,ans2 := 0,0

    for _,num := range nums1{
       if mp2[num] == true {
           ans1++
       }
    } 

     for _,num := range nums2{
       if mp1[num] == true {
           ans2++
       }
    } 

    return append(res, ans1, ans2)
}