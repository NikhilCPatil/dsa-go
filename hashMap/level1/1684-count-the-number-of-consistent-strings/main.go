func countConsistentStrings(allowed string, words []string) int {
    count := 0
    for i:=0;i<len(words);i++{
        for _,char := range words[i] {
            if !strings.ContainsRune(allowed,char){
                count--
                break
            } 
        }
            count++
    }

    return count
}