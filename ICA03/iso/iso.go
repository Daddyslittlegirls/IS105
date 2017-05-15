package main

import (
  "fmt"

)

func main() {
  finalText := makeString()
  IterateOverASCIIStringLiteral(finalText)
}

func makeString() string {
  var byteHolder []byte
  for i := 0x80; i <= 0xff; i++ {
    byteHolder = append(byteHolder, byte(i))
  }
  extendedAscii := string(byteHolder)
  return extendedAscii
}

func IterateOverASCIIStringLiteral(text string) {
	for c := 0; c < len(text); c++ {
    fmt.Printf("%X %c %b \n", text[c], text[c], text[c])
  }
}

// Kode for Oppgave 2b
func GreetingExtendedASCII() {}
