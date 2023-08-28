package guitar

// GuitarType enum
type GuitarType int

const (
	Acoustic GuitarType = iota
	Electric
	Other
)

var guitarTypeStrings = [...]string{
	"Acoustic",
	"Electric",
	"Other",
}

// String method to get the string representation of GuitarType
func (gt GuitarType) String() string {
	if gt < 0 || int(gt) >= len(guitarTypeStrings) {
		return "Unknown"
	}
	return guitarTypeStrings[gt]
}

// WoodType enum
type WoodType int

const (
	Mahogany WoodType = iota
	Maple
	Rosewood
	Cedar
	OtherWood
)

var woodTypeStrings = [...]string{
	"Mahogany",
	"Maple",
	"Rosewood",
	"Cedar",
	"OtherWood",
}

// String method to get the string representation of WoodType
func (wt WoodType) String() string {
	if wt < 0 || int(wt) >= len(woodTypeStrings) {
		return "Unknown"
	}
	return woodTypeStrings[wt]
}

// BuilderType enum
type BuilderType int

const (
	Fender BuilderType = iota
	Gibson
	Martin
	Taylor
	OtherBuilder
)

var builderTypeStrings = [...]string{
	"Fender",
	"Gibson",
	"Martin",
	"Taylor",
	"Other Builder",
}

// String method to get the string representation of BuilderType
func (bt BuilderType) String() string {
	if bt < 0 || int(bt) >= len(builderTypeStrings) {
		return "Unknown"
	}
	return builderTypeStrings[bt]
}
