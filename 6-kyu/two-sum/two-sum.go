package kata
​
func TwoSum(nums []int, target int) [2]int {
    for i := 1 ; i < len(nums); i++{
        for j := i; j < len(nums); j++ {
            if nums[j] + nums[j-i] == target {
                    return [2]int{j-i,j}
            }
        }
    }
    return [2]int {}
}