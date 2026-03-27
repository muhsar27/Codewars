package kata
​
import "strings"
​
func solve(str string) string {
  var lower_count int  
  var upper_count int
  
  for _ , char := range(str){
    if ('a' <= char && char <= 'z'){
      lower_count += 1
       
    }else if ('A' <= char && char <= 'Z'){
      upper_count += 1
    }
  }  
  
  if upper_count > lower_count {
    str = strings.ToUpper(str)
    return str 
  }else if lower_count >= upper_count{
    str = strings.ToLower(str)
    return str 
  }
  return str 
}