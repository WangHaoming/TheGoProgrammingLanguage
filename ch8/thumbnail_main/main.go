package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/WangHaoming/TheGoProgrammingLanguage/ch8/thumbnail"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
