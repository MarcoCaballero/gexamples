package main

// Main examples from UC3M training

import (
	"fmt"
	"os"

	"github.com/marcocab/gexamples/utils"
)

func main() {
	fmt.Printf("Sign (%v) expected (%v)\n", utils.Sign(-10), -1)
	fmt.Printf("Sign (%v) expected (%v)\n", utils.Sign(-11), -1)
	fmt.Printf("Sign (%v) expected (%v)\n", utils.Sign(5), +1)
	fmt.Printf("Sign (%v) expected (%v)\n", utils.Sign(17), +1)
	fmt.Printf("SumOfMultiple (%v) expected (%v)\n", utils.SumOfMultiple(3, 10), 18)
	fmt.Printf("SumOfMultiple (%v) expected (%v)\n", utils.SumOfMultiple(5, 10), 15)
	fmt.Printf("SumOfMultiples3And5 (%v) expected (%v)\n", utils.SumOfMultiples3And5(10), 33)
	fmt.Printf("SumOfMultiples3And5 (%v) expected (%v)\n", utils.SumOfMultiples3And5(50000), 583291668)
	drawer := utils.NewDrawer(2)
	drawer.Draw()
	fmt.Println("\n", len(os.Args), " ", os.Args)
}
