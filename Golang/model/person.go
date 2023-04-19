package model

import "time"

type Person struct{
	Name      string
	Address   Address
	Birthday  time.Time
	Age       int
}

func (p *Person) CalculateAge(){
	year := p.Birthday.Year()
	currYear := time.Now().Year()
	p.Age = (currYear - year)
}

func CalcAge(p Person) int{
	year := p.Birthday.Year()
	currYear := time.Now().Year()
	return (currYear - year)
}