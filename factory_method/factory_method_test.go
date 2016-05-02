package factory_method

import (
	"testing"
)

func TestSingleton(t *testing.T) {

	circlefactory := &CircleFactory{}
	cir := circlefactory.Create(100, 110, 52)
	cir.Draw()

	recfactory := &RactangleFactory{}
	rec := recfactory.Create(100, 110, 52, 102)
	rec.Draw()

}
