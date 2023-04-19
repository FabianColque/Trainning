package model

import "time"

type Compra struct{
	Mercado string
	Data 	time.Time
	Items  	[]ItemsCompra
}

type ItemsCompra struct{
	Name string
}