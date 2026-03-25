package kata
​
import "strings"
​
func GetCount(str string) (count int) {
  for _, char := range str{
    if strings.Contains("aAeEiIoOuU", string(char)){
      count++
    }
  }
  return count
}