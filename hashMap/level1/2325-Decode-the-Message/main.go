func decodeMessage(key string, message string) string {
    var decodedStr strings.Builder
    secKeyMp := make(map[rune]rune)
    i := 'a'
    for _,char := range key{
        if secKeyMp[char] == 0 && unicode.IsSpace(char) != true {
            secKeyMp[char] = i
            i++
        }
    }

     for _,char := range message{
         if unicode.IsSpace(char) == true{
             decodedStr.WriteRune(' ')
         }else{
         decodedStr.WriteRune(secKeyMp[char])
         }
       
    }

    return decodedStr.String()
}