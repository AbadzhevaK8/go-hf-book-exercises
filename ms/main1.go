package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
)

func main() {
	z, _ := zip.OpenReader("task.zip")
	//fmt.Println(z.Name, "z")
	defer z.Close()
	for _, f := range z.File {
		//fmt.Println(f.Name, "f")
		r, _ := f.Open()
		//fmt.Println(r.Name, "r")
		if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
			fmt.Println(f.Name, rows[4][2])
		}
		r.Close()
	}
}
