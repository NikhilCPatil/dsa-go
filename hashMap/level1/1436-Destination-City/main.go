func destCity(paths [][]string) string {
    mp := make(map[string]string)
    for i:=0;i<len(paths);i++{
        mp[paths[i][0]] = paths[i][1]
    }
    fmt.Println(mp)
     
    for _,p := range paths{
        if _,ok:=mp[p[1]]; !ok{
            return p[1]
        }
    }

    return ""
}