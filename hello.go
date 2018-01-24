package main

import ("fmt" //format package
        /*"time"*/)

func main(){
  a := "Applepie"
  fmt.Printf("Meine Variable a ist ein %T\n",a)
}
/*func test1(b,c int) int{
  return b+c
}
func main(){
  fmt.Println(test1(3,4))
}*/
/*func main(){
  for z := 0; z <= 5; z++{
    defer fmt.Println(z)
    }
    }*/

/*func main(){
  time :=time.Now()
  switch{
  case time.Hour()< 12:
    fmt.Println("It`s before noon")
  default:
    fmt.Println("It`s afternoon")
  }*/

/*func main(){
  slice := []int{7,9,10,15}
  fmt.Println("slice==", slice)
  for i := 0; i < len(slice); i++{
    fmt.Printf("slice[%d] == %d\n",i, slice[i])

  slice = append(slice, 11, 22)
  fmt.Printf("slice[%d] == %d\n",i, slice[i])
}}
