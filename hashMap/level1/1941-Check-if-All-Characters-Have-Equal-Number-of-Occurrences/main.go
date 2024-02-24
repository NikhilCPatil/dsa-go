func areOccurrencesEqual(s string) bool {
    mp := make(map[int]int)
    pre := 0
    for _,char := range s{
        pre = int(char - 'a') 
        mp[pre]++
    }

    for _,c := range mp{
        if mp[pre] !=  c {
            return false
        }
    }

    return true
}