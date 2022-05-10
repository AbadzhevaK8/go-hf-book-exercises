package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("task.data")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	lines, _ := r.ReadAll()
	for index, val := range lines[0] {
		if val == "0" {
			fmt.Println("Result:", index+1)
		}
	}
}
