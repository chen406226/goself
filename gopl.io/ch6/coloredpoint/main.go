// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 161.

// Coloredpoint demonstrates struct embedding.
package main

import (
	"fmt"
	"math"
)

//!+decl
import "image/color"

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

//!-decl

func (p Point) Distance(q Point) float64 {
	dX := q.X - p.X
	dY := q.Y - p.Y
	return math.Sqrt(dX*dX + dY*dY)
}
//（p Point）这样不定义指针修改不成功
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
//下面两句添加不添加都是等价的
func (p ColoredPoint) Distance(q Point) float64 {    return p.Point.Distance(q) }
func (p *ColoredPoint) ScaleBy(factor float64) {    p.Point.ScaleBy(factor) }

func main() {
	//!+main
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	// fmt.Println(p.Distance(q)) // 这样会报错compile error: cannot use q (ColoredPoint) as Point
	// 可以看作是继承，ColoredPoint可以继承Point的方法。 但是Distance函数参数是Point类型所以传入ColoredPoint类型不可以
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	//!-main
}

/*
//!+error
	p.Distance(q) // compile error: cannot use q (ColoredPoint) as Point
//!-error
*/

func init() {
	//!+methodexpr
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println("dsfk",p.Distance(q))
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)            // "{2 4}"
	fmt.Printf("%T\n", scale) // "func(*Point, float64)"
	//!-methodexpr
}

func init() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	//!+indirect
	type ColoredPoint struct {
		*Point
		Color color.RGBA
	}

	p := ColoredPoint{&Point{1, 1}, red}
	q := ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
	//!-indirect
}
