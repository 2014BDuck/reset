// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.

// Example 1:

// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0. So it is possible.
// Example 2:

// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.

// Constraints:

// 1 <= numCourses <= 105
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// All the pairs prerequisites[i] are unique.

func canFinish(numCourses int, prerequisites [][]int) bool {
    m := map[int][]int{}
    var need []int
    learned := map[int]bool{}

    for _, v := range prerequisites {
        m[v[0]] = append(m[v[0]], v[1])
    }

    for i := 0; i < numCourses; i++ {
        need = append(need, i)
    }

    for len(m) > 0 {
        canLearn := false

        var newNeed []int
        for _, v := range need {
            var preReq []int
            for _, course := range m[v] {
                if _, ok := learned[course]; !ok {
                    preReq = append(preReq, course)
                }
            }
            if len(preReq) == 0 {
                canLearn = true
                delete(m, v)
                learned[v] = true
            } else {
                m[v] = preReq
                newNeed = append(newNeed, v)
            }
        }
        need = newNeed

        if len(m) > 0 && !canLearn {
            return false
        }
    }
    return true
}