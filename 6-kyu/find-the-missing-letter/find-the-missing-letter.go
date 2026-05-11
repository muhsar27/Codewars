package kata
​
func FindMissingLetter(chars []rune) rune {
​
  for i := 0; i < len(chars)-1; i++ {
    if chars[i+1] != rune(chars[i]+1) {
        return rune(chars[i]+1)
    }
}
​
  return 'a'
}