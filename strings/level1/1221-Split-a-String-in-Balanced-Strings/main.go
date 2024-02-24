func balancedStringSplit(s string) int {
    ans := 0
    balanced := 0
    for _, char := range s{
        if string(char) == "L"{
          balanced++  
        }else{
            balanced--
        }
        if balanced == 0 {
            ans++
        }
    }
    return ans
}