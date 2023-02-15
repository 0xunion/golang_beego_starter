package main

type A interface {
	p()
}

type AImp struct {
	A
}

func (a AImp) p() {}

type Aa struct {
	AImp
}

type Ab struct {
	AImp
}

type B struct {
	o A
}

type C struct {
}

func main() {
	b := B{}
	b.o = Aa{}

	switch b.o.(type) {
	case Aa:
		println("Aa")
	case Ab:
		println("Ab")
	}
}
