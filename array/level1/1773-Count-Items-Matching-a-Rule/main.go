func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0   
	keySet := map[string]int{
		"type" : 0,
		"color" : 1,
		"name" : 2,
	}
   
	   for i:=0;i<len(items);i++{
		   if items[i][keySet[ruleKey]] == ruleValue  {
			   res++
		   }
	   }
   
	   return res
   }