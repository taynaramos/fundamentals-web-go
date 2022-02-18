package main

import (
	"net/http"

	"fundamentals-web-go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
