package main

import "fmt"

var metersPerLiter float64

func paintNeeded(width, height float64) float64 {
	area := width * height
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / metersPerLiter
}

func main() {
	metersPerLiter = 10.0
	fmt.Printf("%.2f\n", paintNeeded(4.2, 3.0))
	/*paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)*/
}
