package main

import "fmt"

func zeroVal(ival int){
    ival = 0
}

func zeroPtr(iptr *int){
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("Initial : ", i)

    zeroVal(i)
    fmt.Println("ZeroVal : ", i)

    zeroPtr(&i)
    fmt.Println("ZeroPtr : ", i)

    fmt.Println("Pointer : ", &i)
}
