func mostWordsFound(sentences []string) int {
    maxLen := 0
    for i:=0;i<len(sentences);i++{
        count := len(strings.Split(sentences[i], " "))
        if count > maxLen{
            maxLen = count
        }
    }
    return maxLen
}