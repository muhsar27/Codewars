package kata
​
​
func MinMax(arr []int) [2]int {
  for i := 0 ; i < len(arr)-1 ; i++{
    for j := 0 ; j < len(arr)-1; j++{
      if (arr[j] > arr[j+1]){
        arr[j],arr[j+1] = arr[j+1],arr[j]
      }
    }
  }
  newArr := [2]int{arr[0], arr[len(arr)-1]}
  return newArr
}
​