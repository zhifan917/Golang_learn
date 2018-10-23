package main

import (
	"fmt"
)

type Developer struct { //开发
	Name string
	Base int
}

func (d *Developer) Calc() int {
	return d.Base
}

type PM struct { //pm
	Name   string
	Base   int
	Option int
}

func (p *PM) Calc() int {
	return p.Base + p.Option
}

type YY struct { //运营
	Name   string
	Base   float32
	Option float32
	Rate   float32 //0.6 ~ 3
}

func (p *YY) Calc() float32 {
	return p.Base + p.Option*p.Rate
}

type EmployeeMgr struct { //最终汇总
	devList []*Developer //用切片存
	pmList  []*PM
	yyList  []*YY
}

func (e *EmployeeMgr) Calc() float32 { //进行计算
	var sum float32
	for _, v := range e.devList { //计算程序员
		sum += float32(v.Calc())
	}

	for _, v := range e.pmList { //计算pm
		sum += float32(v.Calc())
	}

	for _, v := range e.yyList { //计算运营
		sum += float32(v.Calc())
	}

	return sum
}

func (e *EmployeeMgr) AddDev(d *Developer) { //人从哪来，要添加人，添加到列表的函数
	e.devList = append(e.devList, d)
}

func (e *EmployeeMgr) AddPM(d *PM) {
	e.pmList = append(e.pmList, d)
}

func (e *EmployeeMgr) AddYY(d *YY) {
	e.yyList = append(e.yyList, d)
}

func main() {
	var e = &EmployeeMgr{}

	dev := &Developer{ //添加具体人
		Name: "develop",
		Base: 10000,
	}
	e.AddDev(dev)

	pm := &PM{
		Name:   "pm",
		Base:   10000,
		Option: 12000,
	}
	e.AddPM(pm)

	yy := &YY{
		Name:   "yy",
		Base:   10000,
		Option: 12000,
		Rate:   1.2,
	}
	e.AddYY(yy)

	sum := e.Calc() //计算所有人
	fmt.Printf("sum:%f\n", sum)
}
