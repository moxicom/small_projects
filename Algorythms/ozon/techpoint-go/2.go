package main

import "fmt"

func main() {
    var s string
    var t, start_idx, end_idx int
    fmt.Scan(&s)

    byteString := []byte(s)

    fmt.Scan(&t)

    for i := 0; i < t; i++ {
        // fmt.Println()
        fmt.Scanf("%v %v %s", &start_idx, &end_idx, &s)
        start_idx--;
        end_idx--;
        for j := start_idx; j <= end_idx; j++ {
            // fmt.Printf("old byte = %v\n", string(byteString[j]))
            byteString[j] = s[j - start_idx]
            // fmt.Printf("new byte = %v\n", string(byteString[j]))
        }
    }

    fmt.Println(string(byteString))

}
