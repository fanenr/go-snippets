// Dup prints the count and text of lines that appear more than once in the input.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {}

func print(counts map[string]int) {
	fmt.Printf("\ncount\tline\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	print(counts)
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) < 1 {
		count(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v", err)
				continue
			}
			count(f, counts)
			f.Close()
		}
	}
	print(counts)
}

func count(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, file := range os.Args[1:] {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for line := range strings.SplitSeq(string(data), "\n") {
			counts[line]++
		}
	}
	print(counts)
}
