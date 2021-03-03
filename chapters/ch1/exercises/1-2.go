package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Command Line Args")
	for i, arg := range os.Args[1:] {
		io.WriteString(os.Stdout, fmt.Sprintf("%d: %s\n", i, arg))
	}
}
