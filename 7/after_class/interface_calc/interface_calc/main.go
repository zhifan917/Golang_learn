package main

import (
	"fmt"
)

type Employee interface { //定义1个雇员的接口，其规定的方法是calc（计算工资）
	Calc() float32 //接下来的各种职位要想使用这个接口，就需要有该接口规定的方法，并且类型也要一致
}

type Developer struct {
	Name string
	Base float32
}

func (d *Developer) Calc() float32 {
	return d.Base
}

type PM struct {
	Name   string
	Base   float32
	Option float32
}

func (p *PM) Calc() float32 {
	return p.Base + p.Option
}

type YY struct {
	Name   string
	Base   float32
	Option float32
	Rate   float32 //0.6 ~ 3
}

func (p *YY) Calc() float32 {
	return p.Base + p.Option*p.Rate
}

type EmployeeMgr struct {
	employeeList []Employee //员工管理列表不需要在区分职位了，只需要1个职员列表即可（因为无论是dev还是pm、yy都实现了employee的接口）
}

func (e *EmployeeMgr) Calc() float32 { //计算也是都统一一个了
	var sum float32
	for _, v := range e.employeeList {
		sum += v.Calc()
	}

	return sum
}

func (e *EmployeeMgr) AddEmpoyee(d Employee) { //追加也是都只有1个了，并且这里参数也不需要加*，因为interface自身就是引用类型。
	e.employeeList = append(e.employeeList, d)
}

func main() {
	var e = &EmployeeMgr{}

	dev := &Developer{
		Name: "develop",
		Base: 10000,
	}
	e.AddEmpoyee(dev)

	pm := &PM{
		Name:   "pm",
		Base:   10000,
		Option: 12000,
	}
	e.AddEmpoyee(pm)

	yy := &YY{
		Name:   "yy",
		Base:   10000,
		Option: 12000,
		Rate:   1.2,
	}
	e.AddEmpoyee(yy)

	sum := e.Calc()
	fmt.Printf("sum:%f\n", sum)
}
