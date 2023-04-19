package main

import ("fmt"
		"math")

type geometry interface{
	area() float64
}

type rectangle struct{
	width, height float64
}

func (r rectangle) area() float64{
	return r.width * r.height
}

type circle struct{
	radius float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func ShowGeometry(g geometry){
	fmt.Println(g.area())
}

/*func calcRectArea(rect rectangle){
	area := rect.width * rect.height
	fmt.Println(area)
}

func calcCircleArea(circle circle){
	area := math.Pi * circle.radius * circle.radius
	fmt.Println(area)
}*/

func main(){
	fmt.Println("Init...")
	
	rectan := rectangle{
		width : 1,
		height : 2,
	}

	circ := circle{
		radius : 3,
	}

	//calcRectArea(rectan)
	//calcCircleArea(circ)
	ShowGeometry(rectan)
	ShowGeometry(circ)
}