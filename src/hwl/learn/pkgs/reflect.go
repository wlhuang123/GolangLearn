package pkgs

import (
	"hwl/tool"
	"hwl/tool/logs"
	"reflect"
)

// ReflectTest .
func ReflectTest() {
	showVariableType()
	exchange()
	showVariableValue()
	panicIfReflectTypeNotMatch()
	modifyValue()
	getStructField()
	getStructMethod()
}

func showVariableType() {
	var v string
	logs.Println(reflect.TypeOf(v).Kind())
}

func showVariableValue() {
	v := "txt"
	logs.Println(reflect.ValueOf(v))
}

// 变量，interface{}，reflect.Value转化
func exchange() {
	var i interface{}
	s := "string context"
	// 变量转化为intrface{}
	i = s
	// interface{}转化为reflect.Value
	v := reflect.ValueOf(i)
	logs.Println(v.Kind())
	// reflect.Value转化为interface{}
	iFromValue := v.Interface()
	// interface{}转化为变量
	sFromInterface, ok := iFromValue.(string)
	if !ok {
		return
	}
	logs.Println("exchange string context:", sFromInterface)
}

func panicIfReflectTypeNotMatch() {
	tool.PrintFuncHeader("panicIfReflectTypeNotMatch")

	defer func() {
		if err := recover(); err != nil {
			logs.Println(err)
		}
	}()

	var i interface{}
	i = 5
	f := reflect.ValueOf(i).Float()
	logs.Println(f)
}

func modifyValue() {
	tool.PrintFuncHeader("modifyValue")

	var num int
	num = 10

	modify := func(i interface{}) {
		// reflect.ValueOf(i).SetInt(20)  error,因为i接收到的是指针，指针类型不能直接调用SetInt，只有Int类型能调用
		// 需要用Elem方法返回指针或者接口持有的值

		reflect.ValueOf(i).Elem().SetInt(20)
	}

	modify(&num) // 需要传入指针才能对num真正的修改
	logs.Println("num after modify:", num)
}

// Student .
type Student struct {
	Name string `json:"name_student"`
	id   int
}

func getStructField() {
	tool.PrintFuncHeader("getStructField")

	var s Student
	var i interface{}
	i = s
	iValue := reflect.ValueOf(i)
	iType := reflect.TypeOf(i)
	if iValue.Kind() != reflect.Struct {
		logs.Println("expect struct, error")
		return
	}
	num := iValue.NumField()
	logs.Println(num)

	for i := 0; i < num; i++ {
		logs.Println(iType.Field(i).Tag.Get("json"))
	}

}

// Show .
func (s Student) Show() {
	logs.Println("show", s.Name)
}

func (s Student) show() {
	logs.Println("show ", s.Name)
}

// ZShow .
func (s Student) ZShow(text string) {
	logs.Println("show ", text)
}

// 是按照函数名字的ASCII来排序的。小写的比如show无法统计到
func getStructMethod() {
	tool.PrintFuncHeader("getStructMethod")

	s := Student{
		Name: "abc",
	}
	var i interface{}
	i = s
	iValue := reflect.ValueOf(i)
	nums := iValue.NumMethod()
	logs.Println(nums)

	iValue.Method(0).Call(nil)
	iValue.Method(1).Call([]reflect.Value{reflect.ValueOf("text")})
}
