package design

import "hwl/tool/logs"

/*
抽象工厂

好处：
*/

// Luncher 午餐 每一份午餐都需要Cook之后才能吃
type Luncher interface {
	Cook()
}

// Rise 米饭
type Rise struct{}

// Cook .
func (r Rise) Cook() {
	logs.Println("rise cook finish")
}

// Vegetable 蔬菜
type Vegetable struct{}

// Cook .
func (v Vegetable) Cook() {
	logs.Println("vegetable cook finish")
}

// LunchFactory 产生午餐的工厂接口
type LunchFactory interface {
	CreateRice() Luncher
	CreateVegetable() Luncher
}

// SimpleLunchFactory .
type SimpleLunchFactory struct{}

// NewSimpleLunchFactory .
func NewSimpleLunchFactory() SimpleLunchFactory {
	return SimpleLunchFactory{}
}

// CreateRise .
func (SimpleLunchFactory) CreateRise() Luncher {
	return &Rise{}
}

// CreateVegetable .
func (SimpleLunchFactory) CreateVegetable() Luncher {
	return &Vegetable{}
}

// AbstractFactoryTest .
func AbstractFactoryTest() {
	factory := NewSimpleLunchFactory()

	rice := factory.CreateRise()
	rice.Cook()

	vegetable := factory.CreateVegetable()
	vegetable.Cook()
}
