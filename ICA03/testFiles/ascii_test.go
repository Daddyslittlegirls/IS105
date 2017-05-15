package main

import "testing"

// Funksjonen skal generere en utskrift fra en sekvens av bytes,
// dvs. av typen []bytes (det betyr at du må finne den heksadesimale
// eller binære representasjonen av alle tegn i strengen,
// som skal skrives ut (inkludert anførselstegn eller
// “double quotes” på engelsk).
// Funksjonen greetingASCII() returnerer en variabel av typen string,
// som inneholder kun ASCII tegn (ikke utvidet ASCII).
//Gjelder oppgave 1b
func GreetingASCII() string {
	a := []byte("\x22\x48\x65\x6C\x6C\x6F\x20\x3A\x2D\x29\x22")
	return string(a)
}

func TestGreetingASCII(t *testing.T) {
    if !(isASCII(GreetingASCII())){
        t.Fail()
    }
}


func isASCII(text string) bool {
    for _, c := range text {
        if c > 127 {
            return false
        }
    }
    return true
}
