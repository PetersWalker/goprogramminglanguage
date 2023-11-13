// Exercis 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fileNames := make(map[string][]string)

	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n, err")
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}

	fmt.Println("dups\tline\tfiles")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			fmt.Printf(" \t[")
			for _, file := range fileNames[line] {
				fmt.Printf(" %v ", file)
			}
			fmt.Printf("]")

			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if notIn(f.Name(), fileNames[input.Text()]) {
			fileNames[input.Text()] = append(fileNames[input.Text()], f.Name())
		}
	}
}

func notIn(value string, list []string) bool {
	for _, element := range list {
		if element == value {
			return false
		}
	}

	return true
}
