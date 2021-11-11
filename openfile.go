package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//try to avois os.RealAll and io/uitil

	file, _ := os.Open("helloWorld.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var allLines []string
	for scanner.Scan() {
		allLines = append(allLines, scanner.Text())
	}
	file.Close()
	fmt.Println(len(allLines))
	for _, val := range allLines {
		fmt.Println(val)
	}
}
