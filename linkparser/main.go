package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
)

func run(args []string, stdout io.Writer) error {
    if len(args) < 2 {
        return fmt.Errorf("Missing file name to parse")
    }
    fileName := args[1]

    buffer, err := os.ReadFile(fileName)
    if err != nil {
        return fmt.Errorf("Reading the file: %w", err)
    }
    reader := bytes.NewReader(buffer)

    result, err := CollectLinks(reader)
    if err != nil {
        return fmt.Errorf("Collecting links: %w", err)
    }
    for _, link := range result {
        fmt.Fprintf(stdout, "%v => %v\n", link.Href, link.Text)
    }
    return nil
}

func main() {
    if err := run(os.Args, os.Stdout); err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err)
        os.Exit(1)
    }
}
