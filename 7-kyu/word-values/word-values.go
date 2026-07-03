package kata
​
func NameValue(list []string) []int { 
  var sum int
  var arr []int
  for i , words := range list{
    for _ , letters := range words{
      if 'a' <= letters && letters <= 'z'{
        sum += (int(letters) - 96) * (i+1)
      }
    }
    
    arr = append(arr , sum)
    sum = 0 
  }
  return arr
}