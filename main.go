package main

import (
	"fmt"
	"github.com/sraynitjsr/Creational_Patterns"
	"github.com/sraynitjsr/Behaviroal_Patterns"
	"github.com/sraynitjsr/Structural_Patterns"
)

func main() {
	fmt.Println("LLD and Design Patterns in GoLang")
	
	creational.Builder()
	creational.Factory()
	creational.Abstract()
	creational.Singleton()

	behaviroal.Observer()

	structural.Composite()
}
