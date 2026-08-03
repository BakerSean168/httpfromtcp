package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 8)
	currentLine := ""

	for {
		n, err := f.Read(buf)
		if n > 0 {
			text := string(buf[:n])
			parts := strings.Split(text, "\n")
			for i := 0; i < len(parts)-1; i++ {
				fmt.Printf("read: %s\n", currentLine+parts[i])
				currentLine = ""
			}

			currentLine += parts[len(parts)-1]
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

	}

	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
