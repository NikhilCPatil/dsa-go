func toLowerCase(s string) string {
    ret := ""
    for _,char := range s {
        if char >= 'A' && char <= 'Z'{
            char = char + 32
        }
       ret += string(char) 
    }
    return ret
}