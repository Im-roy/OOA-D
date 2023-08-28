package guitar

type GuitarSpec struct {
	builder    BuilderType
	model      string
	guitarType GuitarType
	backWood   WoodType
	topWood    WoodType
}

func (g *GuitarSpec) SetBuilder(builder BuilderType) {
	g.builder = builder
}

func (g *GuitarSpec) SetModel(model string) {
	g.model = model
}

func (g *GuitarSpec) SetType(guitarType GuitarType) {
	g.guitarType = guitarType
}

func (g *GuitarSpec) SetBackWood(backWood WoodType) {
	g.backWood = backWood
}

func (g *GuitarSpec) SetTopWood(topWood WoodType) {
	g.topWood = topWood
}

func (g *GuitarSpec) Builder() BuilderType {
	return g.builder
}

func (g *GuitarSpec) Model() string {
	return g.model
}

func (g *GuitarSpec) Type() GuitarType {
	return g.guitarType
}

func (g *GuitarSpec) BackWood() WoodType {
	return g.backWood
}

func (g *GuitarSpec) TopWood() WoodType {
	return g.topWood
}
