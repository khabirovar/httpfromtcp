package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
		parts := strings.Split(string(buf), "\n")
		for i := 0; i < len(parts)-1; i++ {
			fmt.Printf("read: %s%s\n", line, parts[i])
			line = ""
		}
		line += parts[len(parts)-1]
	}
}
