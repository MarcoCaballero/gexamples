package utils

import (
	"fmt"
	"math"
	"strings"
)

const (
	squareWhiteUnicode    = "⛶"
	squareBlackUnicode    = "⏹"
	verticalBarUnicode    = "│"
	horizontalBarUnicode  = "─"
	upperLeftBarUnicode   = "┌"
	upperRightBarUnicode  = "┐"
	bottomLeftBarUnicode  = "└"
	bottomRightBarUnicode = "┘"
)

type Circle struct {
	radius int
}

type Square struct {
	width int
}

type Drawer struct {
	square Square
	circle Circle
	bucket [][]int
}

func NewDrawer(radius int) *Drawer {
	squareWidth := radius * 2

	buckets := make([][]int, squareWidth)
	for i := range buckets {
		buckets[i] = make([]int, squareWidth)
	}

	return &Drawer{
		square: Square{width: squareWidth},
		circle: Circle{radius: radius},
		bucket: buckets,
	}
}

func (drawer *Drawer) Draw() {
	drawer.drawSquare()
	drawer.drawCircle()
}

func (drawer *Drawer) drawSquare() {
	horizontalLine := strings.Repeat(horizontalBarUnicode, drawer.square.width*3)
	fmt.Println(upperLeftBarUnicode + horizontalLine + upperRightBarUnicode)
	for i := range drawer.bucket {
		fmt.Print(verticalBarUnicode)
		for range drawer.bucket[i] {
			fmt.Print(squareWhiteUnicode)
		}
		fmt.Println(verticalBarUnicode)
	}
	fmt.Println(bottomLeftBarUnicode + horizontalLine + bottomRightBarUnicode)
}

func (drawer *Drawer) drawCircle() {
	// var start float64
	// if drawer.circle.radius%2 == 0 {
	// 	start = 10
	// } else {
	// 	start = 0
	// }

	for angle := float64(10); angle < 360; angle += 10 {
		x := float64(drawer.circle.radius) * math.Cos(angle*math.Pi/180)
		y := float64(drawer.circle.radius) * math.Sin(angle*math.Pi/180)
		// fmt.Println("def: ", y, x)
		fmt.Println("pos: ", toRow(y), toCol(x))
		drawer.bucket[toRow(y)][toCol(x)] = 1
	}
	fmt.Print(drawer.bucket)
}

func toRow(y float64) int {
	switch {
	case y < -1:
		return 3
	case y < 0:
		return 2
	case y < 1:
		return 1
	case y <= 2:
		return 0
	}
	return 100
}

func toCol(x float64) int {
	switch {
	case x < -1:
		return 0
	case x < 0:
		return 1
	case x < 1:
		return 2
	case x <= 2:
		return 3
	}
	return 100
}

// -2, 2
// 0, 3

// 4/2/2 = 1
// 1, n-1 = (1, 3)
// 0.0 0.1 0.2 0.3
// 1.0 1.1 1.2 1.3
// 2.0 2.1 2.2 2.3
// 3.0 3.1 3.2 3.3

// squareWhiteUnicode    = " ⛶ "
// squareBlackUnicode    = " ⏹ "
// ⛶ ⏹ ⏹ ⛶
// ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹
// ⛶ ⏹ ⏹ ⛶

// 0.0 0.1 0.2 0.3 0.4
// 1.0 1.1 1.2 1.3 1.4
// 2.0 2.1 2.2 2.3 2.4
// 3.0 3.1 3.2 3.3 3.4
// 4.0 4.1 4.2 4.3 4.4

// ⛶ ⏹ ⏹ ⏹ ⛶
// ⏹ ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹ ⏹
// ⛶ ⏹ ⏹ ⏹ ⛶

// ⛶ ⛶ ⏹ ⏹ ⛶ ⛶
// ⛶ ⏹ ⏹ ⏹ ⏹ ⛶
// ⏹ ⏹ ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹ ⏹ ⏹
// ⛶ ⏹ ⏹ ⏹ ⏹ ⛶
// ⛶ ⛶ ⏹ ⏹ ⛶ ⛶

// ⛶ ⛶ ⏹ ⏹ ⏹ ⛶ ⛶
// ⛶ ⏹ ⏹ ⏹ ⏹ ⏹ ⛶
// ⏹ ⏹ ⏹ ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹ ⏹ ⏹ ⏹
// ⛶ ⏹ ⏹ ⏹ ⏹ ⏹ ⛶
// ⛶ ⛶ ⏹ ⏹ ⏹ ⛶ ⛶

// 12 = i * r^2
// 12 = pi * 2^2
// 12 = pi * 4
// 12/4 = pi
// 12/4 = 3

// 21 = i * r^2
// 21 = pi * 6.25
// 3.36

//       360|0
//   270       90
// 	  180
// ⛶ ⏹ ⏹ ⛶
// ⏹ ⏹ ⏹ ⏹
// ⏹ ⏹ ⏹ ⏹
// ⛶ ⏹ ⏹ ⛶
