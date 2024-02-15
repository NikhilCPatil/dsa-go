func countPairs(nums []int, target int) int {
	res := 0
 for i := 0;i<len(nums);i++ {
	 for j := i;j<len(nums);j++{
		 if(i == j){
			 continue
		 }
		 if (nums[i]+nums[j]) < target {
			res++
		 }
	 }
 }
 return res
}