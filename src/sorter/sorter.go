package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "infile", "File contains value of sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted value")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile=", *infile, ",outfile=", *outfile, ",algorithm=", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("values=", values)
	} else {
		fmt.Println("Error:", err.Error())
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("File to open the infile ", infile)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			fmt.Println("Error:", err1.Error())
			return
		}
		if isPrefix {
			fmt.Println("A too lang line,seems unexception!")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(outfile string, values []int) (err error) {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create outfile ", outfile)
		return
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
