package asciilines

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "errors"
    "strconv"
)

// Since we are restricted to Ascii, 2-D array of bytes is sufficient
type AsciiLines [][]byte

func LoadTVG(filename string) (*AsciiLines, error) {
    // Declare a single static object to work with
    var state AsciiLines

    // Does file exists?
    _, err := os.Stat(filename)
    if err != nil {
        return nil, err
    }

    // Read entire file
    source, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    // Seperate by lines
    lines := strings.Split(string(source), "\n")

    // Split arguments by spaces on line 1 and check quantity
    arguments := strings.Split(lines[0], " ")
    if len(arguments) != 2 {
        return nil, errors.New("Wrong number of arguments")
    }

    // Read dimensions of array
    xSize, err := strconv.Atoi(arguments[0])
    if err != nil {
        return nil, err
    }

    ySize, err := strconv.Atoi(arguments[1])
    if err != nil {
        return nil, err
    }

    // Check dimensions for sanity
    if xSize < 1 || ySize < 1 {
        return nil, errors.New("Invalid dimension integer value")
    }

    // Initialize 2-D array object
    state = make(AsciiLines, xSize)
    for x := 0; x < xSize; x++ {
        state[x] = make([]byte, ySize)
        for y := 0; y < ySize; y++ {
            state[x][y] = "."[0]
        }
    }

    // Iterate over each line
    for line := 1; line < len(lines); line++ {
        // Only read line if it has stuff in it
        if lines[line] != "" {
            // Split parameters by spaces and check quantity
            parameters := strings.Split(lines[line], " ")
            if len(parameters) != 5 {
                return nil, errors.New("Invalid line length")
            }

            // Read character to write and check for errors
            var character byte
            if len(parameters[0]) == 1 {
                character = parameters[0][0]
            } else {
                return nil, errors.New("Character must be a single ascii character")
            }

            // Read starting position
            xStart, err := strconv.Atoi(parameters[1])
            if err != nil {
                return nil, err
            }

            yStart, err := strconv.Atoi(parameters[2])
            if err != nil {
                return nil, err
            }

            // Read length
            length, err := strconv.Atoi(parameters[4])
            if err != nil {
                return nil, err
            }

            // Read direction as boolean value
            var horizontal bool
            if parameters[3] == "h" {
                horizontal = true
            } else if parameters[3] == "v" {
                horizontal = false
            } else {
                return nil, errors.New("Line must be either horizontal (h) or vertical (v)")
            }

            // Iterate over length
            for i := 0; i < length; i++ {
                // Generate position to write
                var x int
                var y int
                if horizontal {
                    x = xStart
                    y = yStart + i
                } else {
                    x = xStart + i
                    y = yStart
                }

                // If is in bounds, overwrite position with new character
                if x >= 0 && x < len(state) && y >= 0 && y < len(state[0]) {
                    state[x][y] = character
                }
            }
        }
    }

    // Has not early returned, so there must be no errors
    return &state, nil
}

func (a *AsciiLines) Print() {
    // For each row, print it as a string
    for _, line := range *a {
        fmt.Println(string(line))
    }
}