package main

import ("sendEmail/internal/domain/campaign"
		"github.com/go-playground/validator/v10")

func main(){
	campaign := campaign.Campaign{}
	validate := validator.New()
	err 	 := validate.Struct(campaign)

	if err == nil {
		println("Nenhum erro")
	}else{
		validationErros := err.(validator.ValidationErrors)

		for _, v := range validationErros{
			println(v.Error())
		}
	}
}