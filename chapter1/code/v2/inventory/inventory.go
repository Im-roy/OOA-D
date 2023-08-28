package inventory

import (
	"code_v1/guitar"
)

type Inventory struct {
	gutars []guitar.Guitar
}

func (i *Inventory) AddGuitar(serialNumber string, price float64, builder, model, guitarType, blackWood, topWood string) {
	newGuitar := guitar.NewGuitar(serialNumber, price, builder, model, guitarType, blackWood, topWood)
	i.gutars = append(i.gutars, newGuitar)
}

func (i *Inventory) GetGuitar() []guitar.Guitar {
	return i.gutars
}

// search algorithmn
func (i *Inventory) Search(builder, model, guitarType, blackWood, topWood string) []guitar.Guitar {
	resultGuitars := []guitar.Guitar{}
	for _, guitar := range i.gutars {
		if guitar.Builder() == builder && guitar.Model() == model && guitar.Type() == guitarType &&
			guitar.BackWood() == blackWood && guitar.TopWood() == topWood {
			resultGuitars = append(resultGuitars, guitar)
		}
	}
	return resultGuitars
}
