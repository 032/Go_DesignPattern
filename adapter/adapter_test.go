package adapter

import (
	"fmt"
	"testing"
)

func TestAdapterCompo(t *testing.T) {
	var target Target

	adaptee := Adaptee{amount: 140}
	fmt.Printf("amount %v in EUR \n", adaptee.GetEUR())

	target = Adaptor{adaptee}
	fmt.Printf("amount %v in RMB \n", target.GetRMB())
}
