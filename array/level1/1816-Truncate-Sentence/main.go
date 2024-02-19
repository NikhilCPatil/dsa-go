func truncateSentence(s string, k int) string {
	strArr := strings.Split(s, " ") 
	 return strings.Join(strArr[:k], " ")
 }