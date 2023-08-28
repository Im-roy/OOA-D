package guitar

type Guitar struct {
	serialNumber string
	price        float64
	spec         GuitarSpec
}

// constructor for Guitar class
func NewGuitar(serialNumber string, price float64, guitarSpec GuitarSpec) Guitar {
	guitar := Guitar{
		serialNumber: serialNumber,
		price:        price,
		spec:         guitarSpec,
	}
	return guitar
}

// Setter methods
func (g *Guitar) SetSerialNumber(serialNumber string) {
	g.serialNumber = serialNumber
}

func (g *Guitar) SetPrice(price float64) {
	g.price = price
}

func (g *Guitar) SetSpec(spec GuitarSpec) {
	g.spec = spec
}

// Getter methods
func (g *Guitar) SerialNumber() string {
	return g.serialNumber
}

func (g *Guitar) Price() float64 {
	return g.price
}

func (g *Guitar) Spec() GuitarSpec {
	return g.spec
}
