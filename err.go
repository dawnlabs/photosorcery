package main

import (
        "fmt"
        "os"
)

func logAndExit (err error) {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(-1)
}
