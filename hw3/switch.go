package main

import (
    "fmt"
    "runtime"
    "time"
)

func pingPong(rounds int, useSingleThread bool) time.Duration {
    if useSingleThread {
        runtime.GOMAXPROCS(1)
    } else {
        runtime.GOMAXPROCS(runtime.NumCPU())
    }

    ch := make(chan struct{})
    done := make(chan struct{})

    // ping goroutine
    go func() {
        for i := 0; i < rounds; i++ {
            <-ch
            ch <- struct{}{}
        }
        done <- struct{}{}
    }()

    start := time.Now()

    // pong goroutine
    go func() {
        for i := 0; i < rounds; i++ {
            ch <- struct{}{}
            <-ch
        }
        done <- struct{}{}
    }()

    <-done
    <-done

    return time.Since(start)
}

func main() {
    rounds := 1000000

    duration1 := pingPong(rounds, true)  
    duration2 := pingPong(rounds, false) 

    fmt.Println("Single-threaded time:", duration1)
    fmt.Println("Multi-threaded time:", duration2)

    avg1 := duration1.Seconds() / float64(2*rounds)
    avg2 := duration2.Seconds() / float64(2*rounds)

    fmt.Printf("Average switch time (single thread): %.9f s\n", avg1)
    fmt.Printf("Average switch time (multi thread) : %.9f s\n", avg2)
}
