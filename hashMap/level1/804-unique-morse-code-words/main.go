func uniqueMorseRepresentations(words []string) int {
    morseDef := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    mp := map[string]bool{}
    for _,word := range words{
        chars := ""
        for _,w := range word{
            chars = chars + morseDef[int(w-'a')]
        }  
        mp[chars] = true
    }
    return len(mp)
}