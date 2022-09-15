package main

import (
    "fmt"
    "sync"
)

func FuncThatPanics() {
    panic("the function panicked due to some reason")
}

func WrapFuncWithPanicRecovery(f func()) {
    defer func() {
        if r := recover(); r != nil {
            // add additional logging and monitoring...
            fmt.Println("Recovered in f", r)
        }
    }()
    f()
}
func ConcurrentExecutions() {
    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            WrapFuncWithPanicRecovery(FuncThatPanics)
        }()
    }
    wg.Wait()
    fmt.Println("Executed Successfully")
}

func main() {
    fmt.Println("Panic GoRoutines")
    ConcurrentExecutions()
}
