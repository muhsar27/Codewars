package kata
​
import "strings"
​
func toWeirdCase(str string) string {
  var output string 
  str = strings.ToLower(str)
  words := strings.Fields(str)
  
  for _ , word := range words{
    for i:= 0; i < len(word);i++{
      if i % 2 == 0 {
        output += strings.ToUpper(string(word[i]))
}else{
        output += string(word[i])
      }
    }
    output += " "
  }
  output = strings.TrimSuffix(output, " ")
  return output
}