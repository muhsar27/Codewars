package kata
​
func Solution(str string) []string {
  res := []string{}
  for i := 0 ; i < len(str); i += 2{
    if len(str[i:len(str)]) < 2{
      res = append(res, (string(str[i]) + "_"))
    }else{
      res = append(res, str[i:i+2])
    }
  }
  return res 
​
}
​