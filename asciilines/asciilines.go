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

func LoadTVG(filename string) (*AsciiLines, error) {
    var state AsciiLines

    _, err := os.Stat(filename)
    if err != nil {
        return nil, err
    }

    source, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(source), "\n")

    arguments := strings.Split(lines[0], " ")
    if len(arguments) != 2 {
        return nil, errors.New("Wrong number of arguments")
    }

    xSize, err := strconv.Atoi(arguments[0])
    if err != nil {
        return nil, err
    }

    ySize, err := strconv.Atoi(arguments[1])
    if err != nil {
        return nil, err
    }

    if xSize < 1 || ySize < 1 {
        return nil, errors.New("Invalid dimension integer value")
    }

    state = make(AsciiLines, xSize)
    for x := 0; x < xSize; x++ {
        state[x] = make([]byte, ySize)
        for y := 0; y < ySize; y++ {
            state[x][y] = "."[0]
        }
    }

    for line := 1; line < len(lines); line++ {
        if lines[line] != "" {
            parameters := strings.Split(lines[line], " ")
            if len(parameters) != 5 {
                return nil, errors.New("Invalid line length")
            }

            var character byte
            if len(parameters[0]) == 1 {
                character = parameters[0][0]
            } else {
                return nil, errors.New("Character must be a single ascii character")
            }

            xStart, err := strconv.Atoi(parameters[1])
            if err != nil {
                return nil, err
            }

            yStart, err := strconv.Atoi(parameters[2])
            if err != nil {
                return nil, err
            }

            length, err := strconv.Atoi(parameters[4])
            if err != nil {
                return nil, err
            }

            var horizontal bool
            if parameters[3] == "h" {
                horizontal = true
            } else if parameters[3] == "v" {
                horizontal = false
            } else {
                return nil, errors.New("Line must be either horizontal (h) or vertical (v)")
            }

            for i := 0; i < length; i++ {
                var x int
                var y int
                if horizontal {
                    x = xStart
                    y = yStart + i
                } else {
                    x = xStart + i
                    y = yStart
                }

                if x >= 0 && x < len(state) && y >= 0 && y < len(state[0]) {
                    state[x][y] = character
                }
            }
        }
    }

    return &state, nil
}

func (a *AsciiLines) Print() {
    for _, line := range *a {
        fmt.Println(string(line))
    }
}