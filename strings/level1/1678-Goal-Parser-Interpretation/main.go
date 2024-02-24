func interpret(command string) string {
    ret := ""
    i := 0
    for i < len(command){
        // fmt.Println(string(command[i]))
        if string(command[i]) == "G" {
          ret += string(command[i])  
          i++
        }else if string(command[i]) == "(" && string(command[i+1]) == ")" {
            ret += "o" 
            i += 2
        }else {
            ret += "al"
            i +=4
        }
        
    }

    return ret 
}