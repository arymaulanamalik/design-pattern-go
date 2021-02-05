package factory

import (
	"fmt"

	"github.com/arymaulanamalik/design-pattern-go/factory-method/shoes-factory/concrete"
	ifc "github.com/arymaulanamalik/design-pattern-go/factory-method/shoes-factory/interfaces"
)

func GetShoes(typeShoes, name string, size int32) (ifc.Ishoes, error) {
	if typeShoes == "man" {
		return concrete.NewShoesMan(name, size), nil
	} else if typeShoes == "woman" {
		return concrete.NewShoesWoman(name, size), nil
	}

	return nil, fmt.Errorf("Choose type of the shoes")
}
