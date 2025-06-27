// Last updated: 6/28/2025, 3:39:40 AM
func permute(nums []int) [][]int {
    var result [][]int
    var current []int
    used := make([]bool, len(nums))
    backtrack(nums, &current, &result, used)
    return result
}

func backtrack(nums []int, current *[]int, result *[][]int, used []bool) {
    if len(*current) == len(nums) {
        perm := make([]int, len(*current))
        copy(perm, *current)
        *result = append(*result, perm)
        return
    }

    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }
        used[i] = true
        *current = append(*current, nums[i])

        backtrack(nums, current, result, used)

        used[i] = false
        *current = (*current)[:len(*current)-1]
    }
}

