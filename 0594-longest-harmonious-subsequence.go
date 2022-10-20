func findLHS(nums []int) int {
    sort.Ints(nums)
    result := 0
    prevIdx := 0
    fmt.Println(nums)
    for i := range nums {
        if i == prevIdx {
            continue
        }
        if nums[i] - nums[prevIdx] > 1 {
            for nums[i] - nums[prevIdx] > 1 {
                prevIdx++
            }
        }
        if nums[i] - nums[prevIdx] == 1 {
            result = max(result, i-prevIdx+1)
        }
    }

    if nums[len(nums)-1] - nums[prevIdx] == 1 {
        result = max(result, len(nums)-prevIdx)
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
