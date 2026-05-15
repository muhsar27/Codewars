package kata
​
import "math"
​
​
func CloseCompare(a, b, m float64 ) int {
  
  if m >= math.Abs(a-b) {
    return 0
  }
  if a < b {
    return -1 
  }
  if a > b {
    return 1 
  }
  
  
  return 0
}
​