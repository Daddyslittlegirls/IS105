package main

import (
  "fmt"
  "os"
  "bytes"
  "log"
)

func main() {
  fmt.Printf("Filen består av %d linjer", countLines("pg100.txt"))
}

// tar argument filnavn
func countLines(filename string) (int) {
  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }

// Stor buffer pga. filstørrelsel
  buffer := make([]byte, 32*1024)
  counter := 0
  newline := []byte{'\n'}

  // scanner for newline
  for {
    c, err := file.Read(buffer)
    counter += bytes.Count(buffer[:c], newline)

    if err != nil {
      log.Fatal(err)
      } else {
        return counter
      }
    }
}
