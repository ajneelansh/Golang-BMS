package routes

import(
	"github.com/gorilla/mux"
	"github.com/ajneelansh/Golang-BMS/controllers"

)

var RegisterBookRoutes = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.CreateBooks()).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks()).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBooksById()).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook()).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeketeBook()).Methods("DELETE")
}