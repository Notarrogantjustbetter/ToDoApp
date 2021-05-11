package routes

import (
	"net/http"
	
	"github.com/Notarrogantjustbetter/ToDo/v2/database"
	"github.com/Notarrogantjustbetter/ToDo/v2/utils"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.ExecuteTemplate(w, "home.html", nil)
	case "POST":
		r.ParseForm()
		task := r.PostForm.Get("Task")
		database.CreateTask(task)
		http.Redirect(w, r, "/Tasks", http.StatusFound)
	}
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	context, _ := database.GetTasks()
	utils.ExecuteTemplate(w, "tasks.html", context)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	task := r.PostForm.Get("Task")
	database.DeleteTask(task)
	http.Redirect(w, r, "/Tasks", http.StatusFound)
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET", "POST")
	router.HandleFunc("/Tasks", getTasksHandler).Methods("GET")
	router.HandleFunc("/Tasks", deleteTaskHandler).Methods("POST")
	return router
}
