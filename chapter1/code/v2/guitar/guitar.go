package guitar

type Guitar struct {
	serialNumber string
	price        float64
	spec         GuitarSpec
}

// constructor for Guitar class
func NewGuitar(serialNumber string, price float64, builder BuilderType, model string, guitarType GuitarType, blackWood, topWood WoodType) Guitar {
	guitar := Guitar{
		serialNumber: serialNumber,
		price:        price,
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

// Getter methods
func (g *Guitar) SerialNumber() string {
	return g.serialNumber
}

func (g *Guitar) Price() float64 {
	return g.price
}
