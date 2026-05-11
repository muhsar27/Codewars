package kata
​
func GrowingPlant(upSpeed, downSpeed, desiredHeight int) int {
  var height int 
  var count int 
  
    for height < desiredHeight{
      height += upSpeed
      if height >= desiredHeight{
        return count + 1
      }
      height -= downSpeed
        count++
    }  
  return count
}