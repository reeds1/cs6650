package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var m sync.Map
    var wg sync.WaitGroup

    start := time.Now()

    for g := 0; g < 50; g++ {
        wg.Add(1)
        go func(g int) {
            defer wg.Done()
            for i := 0; i < 1000; i++ {
                m.Store(g*1000+i, i) 
            }
        }(g)
    }

    wg.Wait()

    count := 0
    m.Range(func(key, value any) bool {
        count++
        return true
    })

    elapsed := time.Since(start)
    fmt.Println("entries:", count)
    fmt.Println("time taken:", elapsed)
}
