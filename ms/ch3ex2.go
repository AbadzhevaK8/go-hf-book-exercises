package main

import (
	"fmt"
	"log"
	)

func perimeter(x, y float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", x)
	} else if y < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", y)
	} else {
		return x * 2 + y * 2, nil
	}
}

func main() {
	t1, err := perimeter(8.2, 10)
	if err != nil {
		log.Fatal(err)
	}
	t2, err := perimeter(5, 5.4)
	if err != nil {
		log.Fatal(err)
	}
	t3, err := perimeter(6.2, 4.5)
	if err != nil {
		log.Fatal(err)
	}
	total := t1 + t2 + t3
	fmt.Printf("You'll need %.1f meters of fencing\n", total)
}
