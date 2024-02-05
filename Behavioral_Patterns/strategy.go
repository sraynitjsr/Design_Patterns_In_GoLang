package behavioral

import "fmt"

type Strategy interface {
	Execute()
}

type ConcreteStrategyA struct{}

func (a *ConcreteStrategyA) Execute() {
	fmt.Println("Executing ConcreteStrategyA")
}

type ConcreteStrategyB struct{}

func (b *ConcreteStrategyB) Execute() {
	fmt.Println("Executing ConcreteStrategyB")
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	c.strategy.Execute()
}

func Strategy() {
	strategyA := &ConcreteStrategyA{}
	strategyB := &ConcreteStrategyB{}

	context := &Context{strategy: strategyA}

	context.ExecuteStrategy()

	context.SetStrategy(strategyB)

	context.ExecuteStrategy()
}
