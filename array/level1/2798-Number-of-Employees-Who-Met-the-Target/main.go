func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    employesReachedTarget := 0
    for i:=0;i<len(hours);i++{
        if hours[i] >= target {
            employesReachedTarget++
        }
    }
    return employesReachedTarget
}