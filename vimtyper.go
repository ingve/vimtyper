package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	flag.StringVar(&filename, "filename", "test.txt", "Filename")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(`Set FontSize 16
Set Width 1600
Set Height 1000

Sleep 1s
Type "vim"
Sleep 1.5s
Type " %s"
Sleep 500ms
Enter
Type "i"
Enter
`, filename)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Printf("Enter\n")
			continue
		}
		for _, r := range line {
			if r == '\t' {
				fmt.Printf("Tab\n")
				continue
			} else {
				rest := strings.TrimLeft(line, "\t")
				if rest != "" {
					delim := "'"
					if strings.Contains(rest, "'") {
						delim = "`"
					}
					fmt.Printf("Type %s%s%s\nEnter\n", delim, rest, delim)
				}
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(`Sleep 2.5s
Escape
Type ":wq"
Enter
Sleep 2s`)
}
