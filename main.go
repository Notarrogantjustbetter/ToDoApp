package main

import (
	"net/http"

	"github.com/Notarrogantjustbetter/ToDo/v2/database"
	"github.com/Notarrogantjustbetter/ToDo/v2/routes"
	"github.com/Notarrogantjustbetter/ToDo/v2/utils"
)

func main() {
	router := routes.InitRouter()
	utils.LoadTemplate()
	database.InitRedis()
	http.ListenAndServe(":8080", router)
}
