func frequencySort(s string) string {
    m := make(map[byte]int)
    for i := range s {
        m[s[i]]++
    }
    l := make([]string, len(m))
    for k, v := range m {
        str := ""
        for j := 0; j < v; j++ {
            str += string(k)
        }
        l = append(l, str)
    }
    sort.Slice(l, func(i, j int) bool {
        if len(l[i]) > len(l[j]) {
            return true
        }
        return false
    })
    
    return strings.Join(l, "")
}
