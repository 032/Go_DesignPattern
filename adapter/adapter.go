package adapter

type Target interface {
	GetRMB() int // need the rmb amount
}

type Adaptee struct {
	amount int // money amount in euro
}

type Adaptor struct {
	Adaptee
}

func (a Adaptee) GetEUR() int {
	return a.amount
}

// adaptor implements the Target interface
func (a Adaptor) GetRMB() int {
	return a.amount * 7
}
