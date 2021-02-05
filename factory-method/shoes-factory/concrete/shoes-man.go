package concrete

import (
	ifc "github.com/arymaulanamalik/design-pattern-go/factory-method/shoes-factory/interfaces"
)

type ShoesMan struct {
	Shoes
}

func NewShoesMan(name string, size int32) ifc.Ishoes {
	return &ShoesMan{
		Shoes: Shoes{
			Name: name,
			Size: size,
		},
	}
}
