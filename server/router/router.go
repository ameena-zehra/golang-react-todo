package router
// This code sets up the routing configuration for the API endpoints related to tasks in a todo app using the gorilla mux library
// Key Features of Gorilla Mux
// 1. Routing: allows you to define complex url patterns and route requests to specific handlers
// 2. Path Variables: allows you to extract path variables from URLs making it easy to handle dynamic segments of your route
// 3. Middleware: allows it to perform common tasks like authentication and logging (modular reusable code)
// 4. Subroutes: modular hierarchical routes
// 5. HTTP Verbs: provides methods for handling GET, POST, PUT, DELETE, allowing for RESTful API design

import (
	"github.com/ameena-zehra/golang-react-todo/middleware" 
	"github.com/gorilla/mux"
)
func Router() *mux.Router{
	router:= mux.NewRouter()
	router.HandleFunc("/api/tasks", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks" middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return router
}