package main

import (
	"net/http"

	"github.com/taowu99/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
