package creational

import "fmt"

type Product interface {
    GetName() string
}

type ConcreteProductA struct{}

func (c *ConcreteProductA) GetName() string {
    return "Product A"
}

type ConcreteProductB struct{}

func (c *ConcreteProductB) GetName() string {
    return "Product B"
}

type Creator interface {
    CreateProduct() Product
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
    return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
    return &ConcreteProductB{}
}

func Factory() {
    creatorA := &ConcreteCreatorA{}
    productA := creatorA.CreateProduct()
    fmt.Println("Product from creator A:", productA.GetName())

    creatorB := &ConcreteCreatorB{}
    productB := creatorB.CreateProduct()
    fmt.Println("Product from creator B:", productB.GetName())
}
