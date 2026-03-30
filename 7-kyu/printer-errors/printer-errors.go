package kata
​
import "fmt"
​
func PrinterError(s string) string {
  
  count := 0
  for _, char := range s {
    if 'm' < char && char <= 'z'{
        count ++ 
    }
    }
  ans := fmt.Sprintf("%d/%d", count, len(s))
    return ans
  }
​
​