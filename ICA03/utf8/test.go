package main

import (
  "fmt"
)

func main() {

  /* Først ser vi resultatet av bytesekvensen
  */
  fmt.Printf("%s \n", "\xbd\x3f\x3d\x3f\x20\xe2\x8c\x98")

  /* Sjekker om de to tegnene som ikke vises stemmer overens
    ved å sjekke hvilken escaped unicode som vises ved hjelp av %+q verbet.
  */
  fmt.Printf("%+q \n", "\xbd\xe2\x8c\x98")
  fmt.Printf("%+q \n", "½⌘")

  /* Vi ser at "⌘" representeres riktig i bytesekvens men på grunn av
      manglende kapasitet fra terminal fonten vises den ikke.
      "½" viser derimot at den ikke vil la seg vises med hex verdien
      så vi prøver å skrive den direkte som utf8 i stedet -> \xC2\xBD
  */
  fmt.Printf("%s \n", "\xc2\xbd\x3f\x3d\x3f\x20\xe2\x8c\x98")

  /* Dette fungerer bra men vi kan også skrive det som en blanding av unicode
    og hex.
  */
  fmt.Printf("%s \n", "\u00bd\x3f\x3d\x3f\x20\xe2\x8c\x98")

}
