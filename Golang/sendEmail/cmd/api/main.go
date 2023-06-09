package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	//"encoding/json"
	"github.com/go-chi/render"
	"github.com/go-chi/chi/v5/middleware"
)

type product struct{
	ID int
	Name string
}

func main(){
	r := chi.NewRouter()
	r.Use(myMiddleware)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		product := r.URL.Query().Get("product")
		id		:= r.URL.Query().Get("id")
		if product != "" {
			w.Write([]byte(product + id))
		}else{
			w.Write([]byte("Teste Root"))
		}
	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request){
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request){
		//w.Header().Set("Content-Type", "application/json")
		obj := map[string]string{"message" : "success"}
		//b,_ := json.Marshal(obj)
		//w.Write(b)

		render.JSON(w, r, obj)
	})

	r.Post("/product", func(w http.ResponseWriter, r *http.Request){
		var product product
		render.DecodeJSON(r.Body, &product)
		product.ID = 5
		render.JSON(w, r, product)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		println("before")
		next.ServeHTTP(w, r)
		println("after")
	})
}