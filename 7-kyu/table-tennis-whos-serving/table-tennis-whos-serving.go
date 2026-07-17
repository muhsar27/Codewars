package kata
​
func WhoIsServing(currentRound int) int {
 
  ptsPlayed := currentRound - 1
  
  if (ptsPlayed / 2) % 2 == 0{
    return 1
  }
  return 2
}