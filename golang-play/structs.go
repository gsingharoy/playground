package main

import "fmt"

type person struct{
    name string
    age int
    company company
}

type company struct{
    name string
}

func main(){
    fmt.Println(person{"Rahul", 21, company{"Hertha BSC"}})

    fmt.Println(person{name: "Nagal", age: 28})

    fmt.Println(person{name: "Herman"})

    s := person{name: "James", age: 40, company: company{name: "MI6"}}
    fmt.Println(s)

    sp := &s

    sp.age = 42
    sp.company.name = "KGB"
    fmt.Println(sp)
    fmt.Println(s)
}
