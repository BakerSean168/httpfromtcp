package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if n > 0 {
			fmt.Printf("read: %s\n", buf[:n])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

	}
}
