package main

import "fmt"

func scoreSummary(name string, sc1 float64, sc2 float64, sc3 float64) {
	avg := (sc1 + sc2 + sc3) / 3
	fmt.Printf("%10s | %8.2f | %8.2f | %8.2f | %8.2f\n", name, sc1, sc2, sc3, avg)
}

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Score 1", "Score 2", "Score 3", "Average")
	fmt.Println("------------------------------------------------------")
	scoreSummary("Jermaine", 95.4, 82.3, 74.6)
	scoreSummary("Jessie", 79.3, 99.1, 82.5)
	scoreSummary("Lamar", 82.2, 95.4, 77.6)
}
