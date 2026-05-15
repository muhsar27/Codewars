package kata
​
func ReverseSeq(n int) []int {
  res := []int{}
  for i := n ; i > 0 ; i--{
    res = append(res, i)
  }
  return res 
}