func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    res := false
    stringCount1 := strings.Join(word1, "")
    stringCount2 := strings.Join(word2, "")

    if stringCount1 == stringCount2 {
        res = true
    }

    return res
}

