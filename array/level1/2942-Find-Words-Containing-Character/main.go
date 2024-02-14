func findWordsContaining(words []string, x byte) []int {
    res := []int{}
    for i:=0;i<len(words);i++{
        for _,val := range words[i]{
            if val == rune(x) {
                    res = append(res, i)
                    break
            }
        }
    }
    return res
}