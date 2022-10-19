func lexicalOrder(n int) []int {
    result := make([]int, 0, n)
    result = append(result, 1)
    for; len(result) < n; {
        ne := result[len(result)-1] * 10
        for ;ne > n; {
            ne /= 10
            ne++
            for ne % 10 == 0 {
                ne /= 10
            }
        }
        result = append(result, ne)
    }
    return result
}
