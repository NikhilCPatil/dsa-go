func maxWidthOfVerticalArea(points [][]int) int {
	maxDif := 0
	newSlice := make([]int,len(points))
	 for i:=0; i<len(points); i++{
		 newSlice[i] = points[i][0]
	 }
	 sort.Ints(newSlice)
	 
	 for i:=0; i<len(newSlice)-1; i++{
		 temp :=newSlice[i+1] - newSlice[i] 
		 if(temp > maxDif){
			 maxDif = temp
		 }
	 }
 
	return maxDif 
 }