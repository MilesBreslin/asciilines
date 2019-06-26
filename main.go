//usr/bin/go run $0 $@

package main

import (
    "os"
    "fmt"
    "./asciilines"
)

func main() {
    for _, arg := range os.Args[1:] {
        ascii, err := asciilines.LoadTVG(arg)
        if err != nil {
            fmt.Println(err)
        }
        ascii.Print()
    }
}