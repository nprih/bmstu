package main

import (
	"fmt"
	"io"
	"os"
)

type NewWriter struct {
	w   io.Writer
	err error
}

func (nw *NewWriter) WriteNewLine(line string) {
	if nw.err != nil {
		return
	}
	_, nw.err = fmt.Fprintln(nw.w, line)
}
func WriteToFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	nw := NewWriter{
		w: file,
	}
	nw.WriteNewLine("lesson 15 go")
	nw.WriteNewLine("hello")
	nw.WriteNewLine("This")
	nw.WriteNewLine("is")
	nw.WriteNewLine("best")
	nw.WriteNewLine("BMSTU")
	nw.WriteNewLine("Course")
	nw.WriteNewLine("that's all")
	return nw.err
}

func main() {
	err := WriteToFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All is OK!")
}
