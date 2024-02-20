func numJewelsInStones(jewels string, stones string) int {
    mp := make(map[rune]int)
    jwellCount := 0
    for _,char := range jewels{
        mp[char] = 1
    }

    for _,char := range stones{
        if mp[char] == 1{
           jwellCount++ 
        }
    }
    return jwellCount;
}