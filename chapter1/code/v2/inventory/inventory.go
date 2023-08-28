package inventory

import (
	"code_v2/guitar"
)

type Inventory struct {
	gutars []guitar.Guitar
}

func (i *Inventory) AddGuitar(serialNumber string, price float64, builder guitar.BuilderType, model string, guitarType guitar.GuitarType, blackWood, topWood guitar.WoodType) {
	guitarSpec := guitar.NewGuitarSpec(builder, model, guitarType, blackWood, topWood)
	newGuitar := guitar.NewGuitar(serialNumber, price, guitarSpec)
	i.gutars = append(i.gutars, newGuitar)
}

func (i *Inventory) GetGuitar() []guitar.Guitar {
	return i.gutars
}

// search algorithmn
func (i *Inventory) Search(builder guitar.BuilderType, model string, guitarType guitar.GuitarType, blackWood, topWood guitar.WoodType) []guitar.Guitar {
	resultGuitars := []guitar.Guitar{}

	requiredSpec := guitar.NewGuitarSpec(builder, model, guitarType, blackWood, topWood)
	for _, guitar := range i.gutars {
		if guitar.Spec().Equals(requiredSpec) {
			resultGuitars = append(resultGuitars, guitar)
		}
	}
	return resultGuitars
}
