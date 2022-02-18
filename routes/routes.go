package routes

import (
	"fundamentals-web-go/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)        // tela
	http.HandleFunc("/new", controllers.New)       // tela
	http.HandleFunc("/insert", controllers.Insert) // açao
	http.HandleFunc("/delete", controllers.Delete) // açao
	http.HandleFunc("/edit", controllers.Edit)     // tela
	http.HandleFunc("/update", controllers.Update) // açao
}
