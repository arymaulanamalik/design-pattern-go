package concrete

import (
	ifc "github.com/arymaulanamalik/design-pattern-go/factory-method/shoes-factory/interfaces"
)

type ShoesWoman struct {
	Shoes
}

func NewShoesWoman(name string, size int32) ifc.Ishoes {
	return &ShoesWoman{
		Shoes: Shoes{
			Name: name,
			Size: size,
		},
	}
}
