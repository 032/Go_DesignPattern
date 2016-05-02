package bridge

import "fmt"

type draw interface {
	Draw()
}

type Brush struct {
	Size int
}

type Paint struct {
	Color string
}

type Painter struct {
	Brush Brush
	Paint Paint
}

func (p Painter) Draw() {
	fmt.Printf("painter is drawing with %v size brush and %v color paint", p.Brush, p.Paint)
}
