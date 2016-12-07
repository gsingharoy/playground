package main

import "fmt"

// Just to test control structures :D

func main(){
  i := 0
  for(i <= 10){
    fmt.Println(i, oddOrEven(i))
    i += 1
  }
}

func oddOrEven(number int) string{
  if number % 2 == 0{
    return "even"
  } else {
    return "odd"
  }
}
