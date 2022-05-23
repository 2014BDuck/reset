func isValid(s string) bool {
    rm := map[string]string{
        "}":"{",
        ")":"(",
        "]":"[",
    }
    var stack []string
    for _, c := range s {
        if last, ok := rm[string(c)]; ok {
            if len(stack) < 1 || stack[len(stack) - 1] != last {
                return false
            }else{
                stack = stack[:len(stack)-1]
            }
        }else{
            stack = append(stack, string(c))
        }
    }
    
    return len(stack) == 0
}
