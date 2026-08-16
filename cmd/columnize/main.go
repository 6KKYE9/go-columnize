package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"columnize"
)

func main() {
	sep := ","
	gap := 2
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-sep":
			if i+1 < len(args) {
				sep = args[i+1]
				i++
			}
		case "-gap":
			if i+1 < len(args) {
				if n, err := strconv.Atoi(args[i+1]); err == nil {
					gap = n
				}
				i++
			}
		}
	}
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	fmt.Println(columnize.Columnize(lines, sep, gap))
}
