package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	p := Painter{
		Brush: Brush{Size: 10},
		Paint: Paint{Color: "Red"},
	}
	p.Draw()

	p = Painter{
		Brush: Brush{Size: 20},
		Paint: Paint{Color: "Blue"},
	}
	p.Draw()

}
