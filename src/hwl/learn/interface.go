package learn

type annimal interface {
	live()
}

type bird struct{}
type tiger struct{}

func (b bird) live()   {}
func (t *tiger) live() {}

// InplementPointerAndInsTest 接口的指针实现和实例实现
func InplementPointerAndInsTest() {
	var a []annimal
	var b bird
	var t tiger
	a = append(a, b)
	a = append(a, &b)
	a = append(a, &t)
	// a = append(a, t)  //cannot use t (type tiger) as type annimal in append: tiger does not implement annimal (live method has pointer receiver)
}
