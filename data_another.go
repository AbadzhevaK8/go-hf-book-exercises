package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("task.data")
	if err != nil {
		fmt.Println("Error Open File")
	}
	defer file.Close()

	rd := bufio.NewReader(file)

	for text, err, i := "", error(nil), uint(1); err != io.EOF; i++ {
		text, err = rd.ReadString(';')
		if text == "0;" {
			fmt.Println("text =", text, "number of number 0 is ", i)
			break
		}
	}
}
