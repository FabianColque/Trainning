package main

import ("fmt"
		"exercicio/model"
		"time")


func main(){
	var items []model.ItemsCompra
	items = append(items, model.ItemsCompra{Name : "Arroz"})
	items = append(items, model.ItemsCompra{Name : "Feijao"})

	compra := model.Compra{
		Mercado : "mercado Y",
		Data    : time.Now(),
		Items   : items,
	}

	fmt.Println(compra)
}
