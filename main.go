package main

import (
	"net/http"
	"./utils"
	"./database"
	"./server"
)

func main() {
	utils.LoadTemplate()
	database.InitDb()
	router := server.InitServer()
	http.ListenAndServe(":8080", router)
}