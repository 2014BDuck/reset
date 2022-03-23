// Given an integer n, return the least number of perfect square numbers that sum to n.

// A perfect square is an integer that is the square of an integer; in other words, it is the product of some integer with itself. For example, 1, 4, 9, and 16 are perfect squares while 3 and 11 are not.

// Example 1:

// Input: n = 12
// Output: 3
// Explanation: 12 = 4 + 4 + 4.
// Example 2:

// Input: n = 13
// Output: 2
// Explanation: 13 = 4 + 9.

// Constraints:

// 1 <= n <= 104

func numSquares(n int) int {
	if n <= 2 {
		return n
	}

	lst := make([]int, n, n)
	for i := 1; i*i <= n; i++ {
		lst = append(lst, i*i)
	}

	cnt := 0
	toCheck := map[int]bool{n: true}

	for {
		cnt++
		temp := map[int]bool{}
		for x, _ := range toCheck {
			for _, y := range lst {
				if x == y {
					return cnt
				}
				if x < y {
					break
				}
				temp[x-y] = true
			}
		}
		toCheck = temp
	}
	return cnt
}