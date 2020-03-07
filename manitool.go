package main

import (
    "flag"
    "os"
    "bufio"
    "strings"
    "fmt"
    "log"
)

func readLines(filePath, quoteS, quoteD, stringOld, stringNew string) ([]string, error) {
    filePtr, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open a file for reading: %s", err)
    }
    defer filePtr.Close()
    
    lineList := []string{}
    scannerPtr := bufio.NewScanner(filePtr)
    for scannerPtr.Scan() {
        line := scannerPtr.Text()
        temp := strings.Replace(line, "'",  quoteS, -1)
        temp  = strings.Replace(temp, "\"", quoteD, -1)
        temp  = strings.Replace(temp, stringOld, stringNew, -1)
        temp  = strings.Replace(temp, quoteD, "\"", -1)
        temp  = strings.Replace(temp, quoteS, "'",  -1)
        
        lineList = append(lineList, temp)
    }
    return lineList, scannerPtr.Err()
}

func writeLines(filePath string, lineList []string) error {
    filePtr, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to open a file for writing: %s", err)
    }
    defer filePtr.Close()

    writerPtr := bufio.NewWriter(filePtr)
    for _, line := range lineList {
        _, err := fmt.Fprintln(writerPtr, line)
        if err != nil {
            return fmt.Errorf("failed to write a line: %s", err)
        }
    }
    err = writerPtr.Flush()
    if err != nil {
        return fmt.Errorf("failed to flush: %s", err)
    }
    return nil
}

func main() {
    fmt.Println(*sPtr)
    args := flag.Args()
    if len(args) == 0 {
        log.Fatal("ERROR: no option specified")
    }
    
    switch args[0] {
    case "replace":
        if len(args) != 6 {
            log.Fatalf("ERROR: invalid number of parameters for %s: %d", args[0], len(args))
        }
        filePath  := args[1]
        quoteS    := args[2]
        quoteD    := args[3]
        stringOld := args[4]
        stringNew := args[5]
        
        lineList, err := readLines(filePath, quoteS, quoteD, stringOld, stringNew)
        if err != nil {
            log.Fatal("ERROR: ", err)
        }
        
        err = writeLines(filePath, lineList)
        if err != nil {
            log.Fatal("ERROR: ", err)
        }
    default:
        log.Fatalf("ERROR: no match option: %s", args[0])
    }
}