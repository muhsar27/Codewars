package kata
​
import "strings"
​
func SpinWords(str string) string {
  words := strings.Split(str, " ")
​
  
    for i:= 0 ; i < len(words) ; i++{
      if len(words[i]) > 4{
        var rev string 
        for j := len(words[i])-1 ; j >= 0 ; j--{
            rev += string(words[i][j])
          
        }
        words[i] = rev
      }
    }
    return strings.Join(words, " ")
  
}// SpinWords