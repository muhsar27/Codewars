package kata
​
func Multiple3And5(number int) int {
  var count int 
  for i := 0 ; i < number ; i++{
    if i % 5 == 0 || i % 3 == 0 {
      count += i
    }
  }
  return count 
}