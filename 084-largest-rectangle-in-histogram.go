// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

// Example 1:

// Input: heights = [2,1,5,6,2,3]
// Output: 10
// Explanation: The above is a histogram where width of each bar is 1.
// The largest rectangle is shown in the red area, which has an area = 10 units.
// Example 2:

// Input: heights = [2,4]
// Output: 4

// Constraints:

// 1 <= heights.length <= 105
// 0 <= heights[i] <= 104

func largestRectangleArea(height []int) int {
    n := len(height)
    stack := []int{n}
    height = append(height, 0)
    maxArea := 0
    for i := 0; i < n+1; i++ {

        for height[i] < height[stack[len(stack)-1]] {
            h := height[stack[len(stack)-1]]
            stack = stack[:len(stack)-1] // pop
            w := i - 1 - stack[len(stack)-1]
            if w < 0 {
                w = i
            }
            maxArea = max(maxArea, h*w)
        }
        stack = append(stack, i)
    }
    return maxArea
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}