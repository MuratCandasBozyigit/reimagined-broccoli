package routes

import (
	"github.com/MuratCandasBozyigit/reimagined-broccoli/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/", controllers.GetBooksById).Methods("GET")
	router.HandleFunc("/book/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/", controllers.DeleteBook).Methods("DELETE")

}
