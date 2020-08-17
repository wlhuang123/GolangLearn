package design

import "hwl/tool/logs"

/*
工厂模式

好处：提供一个接口（NewRestaulant），传入不同的参数即可创建不同的对象，而无需指定它们具体的类。
*/

// Restaulanter 定义一个饭店的接口，每个饭店都能提供食物的功能
type Restaulanter interface {
	GetFood()
}

// ARestaulant .
type ARestaulant struct{}

// GetFood .
func (a ARestaulant) GetFood() {
	logs.Println("a restaulant get food success")
}

// BRestaulant .
type BRestaulant struct{}

// GetFood .
func (b BRestaulant) GetFood() {
	logs.Println("b restaulant get food success")
}

// NewRestaulant 用户根据饭店名获取饭店
func NewRestaulant(name string) Restaulanter {
	switch name {
	case "a":
		return ARestaulant{}
	case "b":
		return BRestaulant{}
	}
	return nil
}

// FactoryTest .
func FactoryTest() {
	NewRestaulant("a").GetFood()
	NewRestaulant("b").GetFood()
}
