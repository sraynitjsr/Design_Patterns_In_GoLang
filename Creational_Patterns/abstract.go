package creational

import "fmt"

type Shape interface {
	Draw() string
}

type AbstractFactory interface {
	CreateCircle() Shape
	CreateSquare() Shape
}

type ConcreteCircle struct{}

func (c *ConcreteCircle) Draw() string {
	return "Drawing Circle"
}

type ConcreteSquare struct{}

func (s *ConcreteSquare) Draw() string {
	return "Drawing Square"
}

type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateCircle() Shape {
	return &ConcreteCircle{}
}

func (f *ConcreteFactory1) CreateSquare() Shape {
	return &ConcreteSquare{}
}

type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateCircle() Shape {
	return &ConcreteCircle{}
}

func (f *ConcreteFactory2) CreateSquare() Shape {
	return &ConcreteSquare{}
}

func ClientCode(factory AbstractFactory) {
	circle := factory.CreateCircle()
	square := factory.CreateSquare()

	fmt.Println(circle.Draw())
	fmt.Println(square.Draw())
}

func Abstract() {	
	factory1 := &ConcreteFactory1{}
	ClientCode(factory1)

	fmt.Println("---------------------")

	factory2 := &ConcreteFactory2{}
	ClientCode(factory2)
}
