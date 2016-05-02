package abstract_factory

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type Viewport struct {
	X      int
	Y      int
	Width  int
	Height int
}

type Shape interface {
	Draw()
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

type ShapeFactory interface {
	Create(viewport Viewport) Shape
}

type CircleFactory struct{}

func (factory *CircleFactory) Create(v Viewport) Shape {
	return &Circle{
		Location: Point{X: v.X, Y: v.Y},
		Radius:   v.Width,
	}
}

func (c *Circle) Draw() {
	fmt.Printf("[ Circle Location(%d, %d) Radius=%d ] \n", c.Location.X, c.Location.Y, c.Radius)
}

type RectangleFactory struct{}

func (factory *RectangleFactory) Create(v Viewport) Shape {
	return &Rectangle{
		Location: Point{X: v.X, Y: v.Y},
		Width:    v.Width,
		Height:   v.Height,
	}
}
func (c *Rectangle) Draw() {
	fmt.Printf("[ Rectangle Location(%d, %d) Width=%d Height=%d ] \n", c.Location.X, c.Location.Y, c.Width, c.Height)
}
