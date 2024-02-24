func defangIPaddr(address string) string {
    newStrArr :=  strings.Split(address, ".")
    return  strings.Join(newStrArr, "[.]")
}