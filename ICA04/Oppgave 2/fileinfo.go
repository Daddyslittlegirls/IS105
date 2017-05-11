package main

import (
	"fmt"
	"os"
)

func main() {

fileInfo, err := os.Lstat(os.Args[1])
	if err != nil {
		fmt.Println("Can't read file:", os.Args[1])
		panic(err)
}

fmt.Println("Information about a file", fileInfo.Name(), ":")

fmt.Println("Size:", fileInfo.Size())

if fileInfo.Mode()&os.ModeDir == os.ModeDir {
	fmt.Println("Is a directory")
		} else {
	fmt.Println("Is not a directory")
}

if fileInfo.Mode().IsRegular() {
	fmt.Println("Is a regular file")
		} else {
	fmt.Println("Is not a regular file")
}

fmt.Println("Has Unix permission bits:", fileInfo.Mode().Perm())

if fileInfo.Mode()&os.ModeAppend == os.ModeAppend {
	fmt.Println("Is append only")
		} else {
	fmt.Println("Is not append only")
}

if fileInfo.Mode()&os.ModeDevice == os.ModeDevice {
	fmt.Println("Is a device file")
		} else {
	fmt.Println("Is not a device file")
}

if fileInfo.Mode()&os.ModeCharDevice == os.ModeCharDevice {
	fmt.Println("Is a Unix character device")
		} else {
	fmt.Println("Is not a Unix character device")
}



if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
	fmt.Println("File is a symbolic link")
		} else {
	fmt.Println("File is not a symbolic link")
}

 }
