func cellsInRange(s string) []string {
    res := []string{}
    startCol := s[0]
    startRow := int(s[1]-'0')
    endCol := s[3]
    endRow := int(s[4]-'0')

    c := startCol
    for c <= endCol {
        for i:=startRow;i<=endRow;i++{
            res = append(res, string(c) + strconv.Itoa(i))
        }
        c = c + 1 
    }


    return res
}