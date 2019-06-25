package asciilines

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "errors"
    "strconv"
)

type AsciiLines [][]byte

func LoadTVG(filename string) (AsciiLines, error) {
    var state AsciiLines

    _, err := os.Stat(filename)
    if err != nil {
        return state, err
    }

    source, err := ioutil.ReadFile(filename)
    if err != nil {
        return state, err
    }

    lines := strings.Split(string(source), "\n")

    arguments := strings.Split(lines[0], " ")
    if len(arguments) != 2 {
        return state, errors.New("Wrong number of arguments")
    }

    xSize, err := strconv.Atoi(arguments[0])
    if err != nil {
        return state, err
    }

    ySize, err := strconv.Atoi(arguments[1])
    if err != nil {
        return state, err
    }

    if xSize < 1 || ySize < 1 {
        return state, errors.New("Invalid dimension integer value")
    }

    state = make(AsciiLines, ySize)
    for y := 0; y < ySize; y++ {
        state[y] = make([]byte, xSize)
        for x := 0; x < xSize; x++ {
            state[y][x] = "."[0]
        }
    }

    return state, nil
}

func (a *AsciiLines) Print() {
    for _, line := range *a {
        fmt.Println(string(line))
    }
}