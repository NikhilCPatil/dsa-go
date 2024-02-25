func sortSentence(s string) string {
    sArr := strings.Split(s, " ")
    retArr := make([]string, len(sArr))

    for _,val := range sArr{
        a := strings.Split(val, "")
        pos,_ := strconv.ParseInt(a[len(a)-1], 10, 32)
        retArr[pos-1] = strings.Join(a[ :len(a)-1], "")
    }

    return strings.Join(retArr, " ")
}