package kata
​
​
func DigitalRoot(n int) int {
  var digit int
  for {
    for n >= 1 {
​
      digit += n % 10
      n /= 10
​
      
    }
    if digit/10 < 1 {
      break
    }else{
      n = digit
      digit = 0
      continue
    }
  }
  return digit 
}
​