package guitar

type Guitar struct {
	serialNumber string
	price        float64
	builder      string
	model        string
	guitarType   string
	backWood     string
	topWood      string
}

// constructor for Guitar class
func NewGuitar(serialNumber string, price float64, builder, model, guitarType, blackWood, topWood string) Guitar {
	guitar := Guitar{
		serialNumber: serialNumber,
		price:        price,
		builder:      builder,
		model:        model,
		guitarType:   guitarType,
		backWood:     blackWood,
		topWood:      topWood,
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

func (g *Guitar) SetBuilder(builder string) {
	g.builder = builder
}

func (g *Guitar) SetModel(model string) {
	g.model = model
}

func (g *Guitar) SetType(guitarType string) {
	g.guitarType = guitarType
}

func (g *Guitar) SetBackWood(backWood string) {
	g.backWood = backWood
}

func (g *Guitar) SetTopWood(topWood string) {
	g.topWood = topWood
}

// Getter methods
func (g *Guitar) SerialNumber() string {
	return g.serialNumber
}

func (g *Guitar) Price() float64 {
	return g.price
}

func (g *Guitar) Builder() string {
	return g.builder
}

func (g *Guitar) Model() string {
	return g.model
}

func (g *Guitar) Type() string {
	return g.guitarType
}

func (g *Guitar) BackWood() string {
	return g.backWood
}

func (g *Guitar) TopWood() string {
	return g.topWood
}
