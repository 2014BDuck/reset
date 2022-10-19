func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        if points[i][1] < points[j][1] {
            return true
        }
        return false
    })
    
    result := 0
    arrow := 0
    for i := 0; i < len(points);i++ {
        if result == 0 ||  points[i][0] > arrow{
            result++
            arrow = points[i][1]
        }
    }
    return result
}
