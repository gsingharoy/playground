package main

import "fmt"

func main(){
  fmt.Println("1 + 1 = ", 1.0 + 1.0, ". Ain't this sweet.")

  //var someString string
  someString := `
    I am writing a multiline string.
    Ain't this fun too.
  `
  fmt.Println(someString)
  f()
  const xConst string = "Some constant text"
  fmt.Println(xConst)
}

func f(){
  fmt.Println("A new function has been called now.")
}
