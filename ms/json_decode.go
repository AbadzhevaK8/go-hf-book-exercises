package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var items []struct {
		ID uint64 `json:"global_id"`
	}
	f, _ := os.Open("data-20190514T0100.json")
	json.NewDecoder(f).Decode(&items)
	var sum uint64
	for _, item := range items {
		sum += item.ID
	}
	fmt.Println(sum)
}
