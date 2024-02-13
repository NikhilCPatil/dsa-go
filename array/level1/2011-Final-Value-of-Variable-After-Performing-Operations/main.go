func finalValueAfterOperations(operations []string) int {
    ans := 0

    for i := range operations{
        if strings.Contains(operations[i] , "+"){
            ans++
        }else{
            ans--
        }
    }

    return ans
}