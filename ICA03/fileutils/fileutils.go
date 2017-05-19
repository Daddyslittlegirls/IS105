package main

import (
	"io"
	"log"
	"os"
	"fmt"
)

func main() {
	fmt.Printf("%s \n", []byte(fileToByteslice("../files/lang01.wl")))
	fmt.Printf("%s \n", []byte(fileToByteslice("../files/lang02.wl")))
	fmt.Printf("%s \n", []byte(fileToByteslice("../files/lang03.wl")))
	fmt.Printf("%s \n", []byte(editLetter(fileToByteslice("../files/lang03.wl"))))
}

// Sjekker hele byteslicen for hex verdien til "Ø" og erstatter den med "O"
func editLetter(list []byte) []byte {
	for i := 0; i < len(list); i++ {
		if list[i] == 0xF8 {
			list[i] = 0x6F
		}
	}

	return list
}

func fileToByteslice(filename string) []byte {

	// Open file for reading
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	finfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size_of_slice := finfo.Size()

	// The file.Read() function can read a
	// tiny file into a large byte slice,
	// but io.ReadFull() will return an
	// error if the file is smaller than
	// the byte slice
	byteSlice := make([]byte, size_of_slice)

	_, err = io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	return byteSlice

}
