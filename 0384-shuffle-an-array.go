import "math/rand"

type Solution struct {
    nums []int
}


func Constructor(nums []int) Solution {
    return Solution{
        nums: nums,
    }
}


func (this *Solution) Reset() []int {
    return this.nums
}


func (this *Solution) Shuffle() []int {
    tmp := make([]int, len(this.nums),  len(this.nums))
    for i := 0; i < len(this.nums); i++ {
        tmp[i] = this.nums[i]
    }
	rand.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})
    return tmp
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
