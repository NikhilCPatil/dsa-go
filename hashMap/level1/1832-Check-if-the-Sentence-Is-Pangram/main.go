func checkIfPangram(sentence string) bool {
    setMap := make(map[rune]bool)
    for _, char := range sentence {
         setMap[char] = true
    }
    return len(setMap) == 26
}