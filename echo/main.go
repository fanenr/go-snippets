// Echo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	fmt.Println()
	echo2()
	fmt.Println()
	echo3()
}

func echo1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	var s, sep string
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		// fmt.Println(i, " ", arg)
		fmt.Printf("%d %s\n", i, arg)
	}
	fmt.Println(s)

}

func echo3() {
	var s, sep string
	var start time.Time

	// 1
	start = time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	dur1 := time.Since(start).Nanoseconds()
	fmt.Println(s)
	fmt.Printf("cost %dns\n", dur1)

	fmt.Println()

	// 2
	start = time.Now()
	sep = " "
	s = strings.Join(os.Args[1:], sep)
	dur2 := time.Since(start).Nanoseconds()
	fmt.Println(s)
	fmt.Printf("cost %dns\n", dur2)
}
