//usr/bin/go run $0 $@

package main

import (
    "os"
    "fmt"
    "./asciilines"
)

func main() {
    // If using -h or no arguments show help
    if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        fmt.Println("usage: asciilines <tvg files>")
        return
    }

    // For each argument, assume it is a TVG, load it and print it
    for _, arg := range os.Args[1:] {
        ascii, err := asciilines.LoadTVG(arg)
        if err != nil {
            fmt.Println(err)
        }
        ascii.Print()
    }
}