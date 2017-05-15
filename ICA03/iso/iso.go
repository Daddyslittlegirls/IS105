package main

import (
  "fmt"

)

func main() {
/*  Oppgave 2a
finalText := makeString()
IterateOverExtendedASCIIStringLiteral(finalText)
*/
// Oppgave 2b
  fmt.Println(GreetingExtendedASCII())
}

// itererer gjennom og legger til de nødvendige ascci characterene til en
// byte slice. Deretter konverterer og returner den som string.
func makeString() string {
  var byteHolder []byte
  for i := 0x80; i <= 0xff; i++ {
    byteHolder = append(byteHolder, byte(i))
  }
  extendedAscii := string(byteHolder)
  return extendedAscii
}

func IterateOverExtendedASCIIStringLiteral(text string) {
	for c := 0; c < len(text); c++ {
    fmt.Printf("%X %c %b \n", text[c], text[c], text[c])
  }
}

// Kode for Oppgave 2b
func GreetingExtendedASCII() string {
	a:= []byte("\x53\x61\x6c\x75\x74\x2c\x20\xc3\xa7\x61\x20\x76\x61" +
		"\x20\xc2\xb0\x2d\x29\x20\xe2\x80\x8b\xe2\x82\xac\x35\x30")
	return string(a)
}
