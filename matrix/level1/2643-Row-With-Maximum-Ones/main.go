func rowAndMaximumOnes(mat [][]int) []int {
	count, index := 0, 0    
	 for i := range mat {
		 localCount := 0
		 for j:=0;j<len(mat[i]);j++{
			 if mat[i][j] == 1{
			   localCount++  
			 }
		 }
		 if localCount > count{
			 count = localCount
			 index = i
		 }
	 }
 
	 return []int{index, count}
 }