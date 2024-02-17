func createTargetArray(nums []int, index []int) []int {
    target := make([]int, 0 ,len(nums))
    for i,indx := range index{
        if indx < len(target) {
			target = append(target[:indx+1], target[indx:]...)
            target[indx] = nums[i]
        }else{
            target = append(target, nums[i])
        }
    }
    return target
}