package kata
​
import "strings"
​
func ToCamelCase(s string) string {
  
  words := strings.FieldsFunc(s, func(c rune) bool {
    return c == '_' || c == '-'
  })
  
  for i := range words{
    if i != 0 {
      words[i] = strings.Title(words[i])
    }
  }
  return strings.Join(words, "")
}