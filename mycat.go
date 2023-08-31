package main

import (
	"bufio"
	"flag" //
	"fmt"
	"os"
)

func readFile(filePath string, showLineNumbers bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		// read the line
		line := scanner.Text()
		if showLineNumbers {
			fmt.Printf("%d: %s\n", lineNumber, line)
			lineNumber++
		} else {
			fmt.Println(line)
		}
	}

	return scanner.Err() //正常にファイルの末尾まで読み込まれたらnil
}

func main() {
	// Parse command line arguments
	showLineNumbers := flag.Bool("n", false, "Show line numbers")
	//bool型にすると、-n オプションが指定された時trueが設定される
	//showLineNumbersはポインタ
	flag.Parse()
	filePaths := flag.Args()

	for _, filePath := range filePaths {
		err := readFile(filePath, *showLineNumbers)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file", err)
		}
	}

}
