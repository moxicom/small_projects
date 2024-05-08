package main

import "fmt"

func main() {
    var a, b, t int
    fmt.Scan(&t)

    for i := 0; i < t; i++ {
        fmt.Scanf("%v %v", &a, &b)
        fmt.Printf("%v\n", a + b)
    }
}
