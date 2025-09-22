package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var ops uint64
    var wg sync.WaitGroup
    var mu sync.Mutex

    start := time.Now() 

    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                mu.Lock()
                ops++
                mu.Unlock()
            }
        }()
    }

    wg.Wait()

    elapsed := time.Since(start) 
    fmt.Println("ops:", ops)
    fmt.Println("time taken:", elapsed)
}
