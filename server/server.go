package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"../utils"
	"../database"
)

type Server struct {
	router *mux.Router
}

func InitServer() *mux.Router {
	s := &Server{
		router: mux.NewRouter(),
	}
	return s.router
}

func (s Server) routes() {
	s.router.HandleFunc("/", homeHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/", addTaskHandler().ServeHTTP).Methods("POST")
	s.router.HandleFunc("/Tasks", getTasksHandler().ServeHTTP).Methods("GET")
	s.router.HandleFunc("/Tasks", deleteTaskHandler().ServeHTTP).Methods("POST")
}

func homeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.ExecuteTemplate(w, "home.html", nil)
	}
}

func addTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		task := r.PostForm.Get("Task")
		database.Task{}.CreateTask(task)
		http.Redirect(w, r, "/Tasks", http.StatusFound)
	}
}

func getTasksHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context, _ := database.Task{}.GetTasks()
		utils.ExecuteTemplate(w, "tasks.html", context)
	}
}

func deleteTaskHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		task := r.PostForm.Get("Task")
		database.Task{}.DeleteTask(task)
		http.Redirect(w, r, "/Tasks", http.StatusFound)
	}
}