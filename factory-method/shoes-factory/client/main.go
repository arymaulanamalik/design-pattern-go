package main

import (
	"log"

	"github.com/arymaulanamalik/design-pattern-go/factory-method/shoes-factory/factory"
)

func main() {
	shoes, err := factory.GetShoes("wom", "nike", 32)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(shoes)
}
