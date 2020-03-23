package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

	// input := make([]byte, 20)

	// s, err := reader.Read(input)
	// if err != nil {
	// 	log.Fatalln("Unable to read data")
	// }

	// fmt.Printf("Read %d bytes from stdin\n", s)

	// // Use writer to write output
	// s, err = writer.Write(input)
	// if err != nil {
	// 	log.Fatalln("Unable to write data")
	// }

	// fmt.Printf("Write %d bytes to stdout\n", s)
}
