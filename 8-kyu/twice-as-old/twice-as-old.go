package kata
​
import "math"
​
func TwiceAsOld(dadYearsOld, sonYearsOld int) int { 
  doubleAge := sonYearsOld * 2  
  
  
  
  
  return int(math.Abs(float64(dadYearsOld - doubleAge)));
}