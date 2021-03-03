// Measure time difference between echo1, echo2, and echo3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	printBenchData(bench(echo1))
	printBenchData(bench(echo2))
	printBenchData(bench(echo3))
}

type echoFunc func() string

func bench(f echoFunc) (int64, string) {
	s := time.Now()
	args := f()
	return time.Since(s).Nanoseconds(), args
}

func printBenchData(t int64, args string) {
	fmt.Println("echo1")
	fmt.Println("time: ", t)
	fmt.Println("args: ", args)
	fmt.Println()
}

func echo1() string {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	s := strings.Join(os.Args[1:], " ")
	return s
}
