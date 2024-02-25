func reverseWords(s string) string {
    retStatement := []string{}
    sArr := strings.Split(s, " ")

    for _,val := range sArr {
        revStr := ""
        for i:=len(val)-1;i>=0;i--{
            revStr += string(val[i])
        }
        retStatement = append(retStatement, revStr)
    }

    return strings.Join(retStatement, " ")
}