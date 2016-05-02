package factory_method

import (
	"fmt"
)

type Shape interface {
	Draw()
}

type Point struct {
	X int
	Y int
}

type Circle struct {
	Location Point
	Radius   int
}

type Rectangle struct {
	Location Point
	Width    int
	Height   int
}

type CircleFactory struct{}

func (factory *CircleFactory) Create(x, y, r int) Shape {
	return &Circle{
		Location: Point{X: x, Y: y},
		Radius:   r,
	}
}

func (c *Circle) Draw() {
	fmt.Printf("[ Circle Location(%d, %d) Radius=%d ] \n", c.Location.X, c.Location.Y, c.Radius)
}

type RactangleFactory struct{}

func (factory *RactangleFactory) Create(x, y, width, height int) Shape {
	return &Rectangle{
		Location: Point{X: x, Y: y},
		Width:    width,
		Height:   height,
	}
}
func (c *Rectangle) Draw() {
	fmt.Printf("[ Rectangle Location(%d, %d) Width=%d Height=%d ] \n", c.Location.X, c.Location.Y, c.Width, c.Height)
}
