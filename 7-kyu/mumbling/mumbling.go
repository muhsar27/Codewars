package kata
​
func Accum(s string) string {
    var output string 
  
    for i, char := range s {
      if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z'){
        for j := 0; j <= i; j++ {
          if j == 0 && ('a' <= char && char <= 'z'){
            output += string(char-32)
            continue
          }else if j != 0 && ('A' <= char && char <= 'Z'){
            output += string(char+32)
            continue
          }
          output += string(char)               
        }
        if i < len(s)-1{
          output += "-"
          }
      }
    }
  return output 
}