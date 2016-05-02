package abstract_factory

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	var factory ShapeFactory

	factory = &CircleFactory{}
	viewport := Viewport{
		X:      100,
		Y:      100,
		Width:  55,
		Height: 55,
	}
	cir := factory.Create(viewport)
	cir.Draw()

	factory = &RectangleFactory{}
	viewport = Viewport{
		X:      100,
		Y:      100,
		Width:  55,
		Height: 100,
	}
	rec := factory.Create(viewport)
	rec.Draw()

}
