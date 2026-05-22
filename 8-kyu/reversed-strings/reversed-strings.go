package kata
​
func Solution(word string) string {
  var rev string 
    for i := len(word)-1 ; i >= 0 ; i-- {
        rev += string(word[i])
    }
  return rev 
}