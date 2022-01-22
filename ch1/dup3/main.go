package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	files := make(map[string]map[string]int)
	for _, filename := range os.Args[1:] {
		counts := make(map[string]int)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		fmt.Fprintf(os.Stdin, "data type %T\n", counts)
		// fmt.Fprintf(os.Stdin, "data: %s\n", data)
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
		files[filename] = counts
	}
	for filename, counts := range files {
		fmt.Printf("Filename: %s\n", filename)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}

}
