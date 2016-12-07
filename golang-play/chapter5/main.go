package main

import "fmt"

// Just playing around with arrays, slices and maps

func main() {
  var x [10]int

  x[5] = 254

  fmt.Println(x)

  sliceVar:= make([]int, 10)

  sliceVar[0] = 1
  fmt.Println(sliceVar)

  varMap := make(map[string]int)
  varMap["first"] = 1
  varMap["second"] = 2

  fmt.Println(varMap)
}
