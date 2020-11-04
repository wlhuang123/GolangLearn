package design

import "fmt"

/*
访问者模式

例子：golang和python这两门语言结构体不需要改变，后续需要对语言进行什么功能，直接外部实现接口
比如添加对语言的compile和run，只需要compile和run实现接口即可，golang和python的结构体不需要改变
*/

type visiter interface {
	visitGolang(*golang)
	visitPython(*python)
}

type golang struct {
	language string
}

func (g *golang) accept(v visiter) {
	v.visitGolang(g)
}

type python struct {
	language string
}

func (p *python) accept(v visiter) {
	v.visitPython(p)
}

// VisitorTest .
func VisitorTest() {
	golangInst := golang{"golang"}
	pythonInst := python{"python"}

	golangInst.accept(new(compile))
	pythonInst.accept(new(compile))

	golangInst.accept(new(run))
	golangInst.accept(new(run))
}

type compile struct{}

func (c *compile) visitGolang(g *golang) {
	fmt.Println("compile ", g.language)
}

func (c *compile) visitPython(p *python) {
	fmt.Println("compile ", p.language)
}

type run struct{}

func (r *run) visitGolang(g *golang) {
	fmt.Println("run ", g.language)
}

func (r *run) visitPython(p *python) {
	fmt.Println("run ", p.language)
}
