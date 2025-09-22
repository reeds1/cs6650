package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    const n = 100000

    // -------------------
    // Unbuffered 写入
    // -------------------
    f1, _ := os.Create("unbuffered.txt")
    start := time.Now()
    for i := 0; i < n; i++ {
        f1.Write([]byte(fmt.Sprintf("Line %d\n", i)))
    }
    f1.Close()
    unbufferedTime := time.Since(start)

    // -------------------
    // Buffered 写入
    // -------------------
    f2, _ := os.Create("buffered.txt")
    w := bufio.NewWriter(f2)
    start = time.Now()
    for i := 0; i < n; i++ {
        w.WriteString(fmt.Sprintf("Line %d\n", i))
    }
    w.Flush()
    f2.Close()
    bufferedTime := time.Since(start)

    // -------------------
    // 打印结果
    // -------------------
    fmt.Println("Unbuffered time:", unbufferedTime)
    fmt.Println("Buffered time:", bufferedTime)
}
