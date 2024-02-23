func sortPeople(names []string, heights []int) []string {
    mp := make(map[int]string)
    res := make([]string, 0 ,len(names))
    for i,h := range heights{
        mp[h] = names[i] 
    }
    sort.Ints(heights)
   for i:=len(heights)-1;i>=0;i--{
       res = append(res, mp[heights[i]])
   }

    return res
}