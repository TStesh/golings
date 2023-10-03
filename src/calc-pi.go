// go build -ldflags "-s -w" calc-pi.go
// Calc PI monte-carlo in concurrency mode
// 4 goroutins performs output for ~3,3 sec
package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func calc(n int, c chan int) {
    count := 0
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i := 0; i < n; i++ {
        x := r.Float64()
        y := r.Float64()
        if x*x + y*y <= 1.0 {
            count += 1
        }
    }
    c <- count
}

func main() {
    num := 1_000_000_000
    ch_num := 4
    
    t0 := time.Now()
    
    c := make(chan int, ch_num)
    
    s := strings.Repeat("0", ch_num)
    num4 := num / ch_num
    
    for _ = range s { go calc(num4, c) }
  
    sum := 0
    for _ = range s { sum += <-c }
    
    pi := float64(sum) * 4.0 / float64(num)
    
    t1 := time.Now()
    
    fmt.Printf("pi = %f\n", pi)
    fmt.Printf("Duration %v\n", t1.Sub(t0))
}