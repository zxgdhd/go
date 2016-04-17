// GoGram project GoGram.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//	acceptLineFromStdio()
	//	readTextFromFile()
	readTotalFile()
}

func acceptLineFromStdio() {
	contents := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line != "end" {
			contents[line]++
		} else {
			break
		}
	}
	fmt.Printf("%+v\n", contents)
	for line, n := range contents {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func readTextFromFile() {
	files := os.Args[1:]
	var contents map[string]int
	if len(files) == 0 {
		fmt.Println("请指定文件！")
	} else {
		contents = make(map[string]int)
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%+v\n", err)
				continue
			}
			readContents(file, contents)
			file.Close()
		}
		for line, n := range contents {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func readContents(f *os.File, cont map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		cont[input.Text()]++
	}
}

func readTotalFile() {
	files := os.Args[1:]
	if len(files) != 0 {
		contents := make(map[string]int)
		for _, file := range files {
			data, err := ioutil.ReadFile(file)
			if err == nil {
				content := strings.Split(string(data), "\n")
				for lineNum, lineCont := range content {
					contents[lineCont] = lineNum
				}
			} else {
				fmt.Fprintf(os.Stderr, "dup3:%+v\n", err)
			}
			fmt.Printf("%+v\n", contents)
		}
	} else {
		fmt.Println("没有选择文件！")
	}

}
