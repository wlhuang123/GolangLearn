package design

import "hwl/tool/logs"

/*
观察者模式

案例：当注册一个用户的时候，执行A
不用观察者模式的实现：在注册接口里面执行A，违背单一职责原则。当注册用户又要有B,C,D的事情扩展的时候，会让注册接口变得臃肿。
*/

type observer interface {
	notifyHandle()
}

type aBobserve struct{}

func (a aBobserve) notifyHandle() {
	logs.Println("notify a do something")
}

type bBobserve struct{}

func (b bBobserve) notifyHandle() {
	logs.Println("notify b do something")
}

type regist struct {
	observerList []observer
}

func (r regist) regist(name string) {
	logs.Println(name, " regist succ")
	for _, v := range r.observerList {
		v.notifyHandle()
	}
}

func (r *regist) sub(elems observer) {
	r.observerList = append(r.observerList, elems)
}

// ObserverTest .
func ObserverTest() {
	var r regist
	var a aBobserve
	var b bBobserve
	r.sub(a)
	r.sub(b)

	r.regist("Lisa")
	r.regist("Lilei")
}
