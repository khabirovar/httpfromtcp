package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const FILENAME = "messages.txt"

func main() {
	file, err := os.Open(FILENAME)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 8)

	for {
		count, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
      fmt.Printf("eror: %s\n", err.Error())
		}
		fmt.Printf("read: %v\n", string(buf[:count]))
	}
}
