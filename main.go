package main

import (
	"bytes"
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
	line := ""

	for {
		count, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("eror: %s\n", err.Error())
		}
    buf = buf[:count]
		parts := bytes.Split(buf, []byte("\n"))
    if len(parts) > 1 {
      line += string(parts[0])
      fmt.Printf("read: %s\n", line)
      line = string(parts[1])
    } else {
      line += string(parts[0])
    }
	}
}
