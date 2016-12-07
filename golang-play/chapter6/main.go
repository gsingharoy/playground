package main

import "fmt"
import "sort"

// Play with functions

func main() {
  x:= make([]int, 5)
  x[1] = 100
  x[2] = 30
  x[0] = 10
  x[4] = 18
  average, median := averageAndMedian(x)

  fmt.Println("average is ", average)
  fmt.Println("median is ", median)
}

func averageAndMedian(inputArr []int) (float64, float64) {

  fmt.Println("Calculating average for ... ", inputArr)

  total := 0.0
  for _, element := range inputArr {
    total += float64(element)
  }
  length := float64(len(inputArr))
  average := (total / length)

  sort.Ints(inputArr)

  var median float64

  l := len(inputArr)
  if l % 2 == 0 {
    median = (float64(inputArr[(l / 2) - 1]) + float64(inputArr[(l/2)])) / 2.0
  } else {
    median = float64(inputArr[(l/2)])
  }

  return average, median
}
