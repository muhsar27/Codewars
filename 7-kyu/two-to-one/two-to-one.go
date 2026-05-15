package kata
​
//import "unicode"
import "strings"
​
func TwoToOne(s1 string, s2 string) string {
  concat := s1 + s2 
  var new string 
  runes:= []rune(concat)
  
  for i := 0 ; i < len(runes); i++ {
    for j:= 0 ; j < len(runes)-1; j++{
      if runes[j] > runes[j+1]  {
        runes[j],runes[j+1] = runes[j+1],runes[j]
      }
    }
  }
  for _ , ch := range runes {
    if !strings.ContainsAny(new,string(ch)){
      new += string(ch)
    }
  }
  return new 
}
​