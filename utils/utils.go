package utils // import gexamples/utils

// utils package

// Returns the sign of the argument
func Sign(a int) int {
	if a < 0 {
		return -1
	}

	return 1
}

// Given n calculates the sum of every multiple of 3 and 5 that are no multiple of 15 within the range [0, n].
// So that, the sum of  0, 3, 5, 6, 9, 10, 12, (not 15), 18, 20, 21, 24, 25, 27, (not 30), 33, 35,... 2.999.997, (3.000.000 no)
func SumOfMultiples3And5(n int64) int64 {
	return SumOfMultiple(3, n) + SumOfMultiple(5, n) - SumOfMultiple(15, n)
	// var res int64
	// for i := int64(0); i < n; i++ {
	// 	if i%15 == 0 {
	// 		continue
	// 	}
	// 	if i%3 == 0 || i%5 == 0 {
	// 		res += i
	// 	}
	// }
	// return res
}

// Sum the multiple of d in the interval [0, x]
func SumOfMultiple(d, x int64) int64 {
	// b := d * ((x - 1) / d)
	// n := 1 + ((b - d) / d)
	// return (d + b) * n / 2

	// a primer valor, b último, n apariciones
	// ((a+b) * n)/2
	// d número del que buscamos múltiplos
	// a + d(n-1)
	n := x / d
	a := d
	b := (a + (d * (n - 1)))
	return ((a + b) * n) / 2
}
