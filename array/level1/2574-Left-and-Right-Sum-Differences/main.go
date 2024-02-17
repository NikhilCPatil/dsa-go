func leftRightDifference(nums []int) []int {
    answer := make([]int, len(nums))

    for i:=0;i<len(answer);i++{
        leftSum := getSumOfIndex(nums[:i])
        rightSum := getSumOfIndex(nums[i+1:])
        if leftSum < rightSum {
            temp:= leftSum
            leftSum = rightSum
            rightSum = temp
        }
        answer[i] = leftSum  - rightSum
    }
    return answer
}

func getSumOfIndex(imp []int)int{
    sum := 0
      for i:=0;i<len(imp);i++{
       sum += imp[i]
    }
    return sum 
}