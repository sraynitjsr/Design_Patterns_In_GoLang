package main

import "fmt"

type Product struct {
    Part1 string
    Part2 string
    Part3 string
}

type Builder interface {
    BuildPart1()
    BuildPart2()
    BuildPart3()
    GetProduct() Product
}

type ConcreteBuilder struct {
    product Product
}

func (b *ConcreteBuilder) BuildPart1() {
    b.product.Part1 = "Part 1 built"
}

func (b *ConcreteBuilder) BuildPart2() {
    b.product.Part2 = "Part 2 built"
}

func (b *ConcreteBuilder) BuildPart3() {
    b.product.Part3 = "Part 3 built"
}

func (b *ConcreteBuilder) GetProduct() Product {
    return b.product
}

type Director struct {
    builder Builder
}

func (d *Director) Construct() Product {
    d.builder.BuildPart1()
    d.builder.BuildPart2()
    d.builder.BuildPart3()
    return d.builder.GetProduct()
}

func Builder() {
    builder := &ConcreteBuilder{}

    director := &Director{builder}

    product := director.Construct()

    finalProduct := builder.GetProduct()

    fmt.Println("Product built:", finalProduct)
}
