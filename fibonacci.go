package main

import "fmt"
//calculation
func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    n := 10
    fmt.Printf("Fibonacci sequence up to %d:\n", n)
    for i := 0; i <= n; i++ {
        fmt.Printf("%d ", fibonacci(i))
    }
    fmt.Println()
}
