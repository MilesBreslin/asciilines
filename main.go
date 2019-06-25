//usr/bin/go run $0 $@

package main

import (
    "os"
    "fmt"
    "./pkg/asciilines"
)

func main() {
    for arg, _ := range os.Args {
        fmt.Println(arg)
        ascii, err := asciilines.LoadTVG(os.Args[arg])
        if err != nil {
            fmt.Println(err)
        }
        ascii.Print()
    }
}