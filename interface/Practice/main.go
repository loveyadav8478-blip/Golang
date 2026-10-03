package main

type employee interface {
	getName() string
	getSalary() int
}

type fullTime struct {
	name   string
	salary int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getMessage() string {
	return c.name
}
func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (ft fullTime) getSalary() int {
	return ft.salary
}
func (ft fullTime) getName() string {
	return ft.name
}
