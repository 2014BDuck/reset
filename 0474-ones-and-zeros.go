func findWords(words []string) []string {
    r1, r2, r3 := map[byte]bool{
        'q': true,'w': true,'e': true,'r': true,'t': true,'y': true,'u': true,'i': true,'o': true,'p': true,
    }, map[byte]bool{
        'a': true,'s': true,'d': true,'f': true,'g': true,'h': true,'j': true,'k': true,'l': true,
    }, map[byte]bool{
        'z': true,'x': true,'c': true,'v': true,'b': true,'n': true,'m': true,
    }
    
    result := []string{}
    for i := range words {
        var r map[byte]bool
        flag := false
        if r1[strings.ToLower(words[i])[0]] {
            r = r1
        }else if r2[strings.ToLower(words[i])[0]] {
            r = r2
        }else{
            r = r3
        }
        for j := range words[i] {
            if !r[strings.ToLower(words[i])[j]] {
                flag = true
                break
            }
        }
        if !flag {
            result = append(result, words[i])
        }
    }
    return result
}
    
