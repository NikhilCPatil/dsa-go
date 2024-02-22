func countPoints(rings string) int {
	d := ['Z']int{'R': 1, 'G': 2, 'B': 4}
	ans := 0
	 mask := [10]int{}
	 for i, n := 0, len(rings); i < n; i += 2 {
		 c := rings[i]
		 j := int(rings[i+1] - '0')
		 mask[j] |= d[c]
	 }
	 fmt.Println(mask)
	 for _, x := range mask {
		 if x == 7 {
			 ans++
		 }
	 }
	 return ans
 }