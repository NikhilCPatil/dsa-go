func restoreString(s string, indices []int) string {
    arr := make([]string, len(indices))
    for i:=0;i<len(indices);i++{
        arr[indices[i]] = string(s[i]) 
    }
    return strings.Join(arr, "")
}