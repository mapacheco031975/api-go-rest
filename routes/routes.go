package routes

import (
	"api-go-rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/api/personalidades/{id}", controllers.AlterarPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades", controllers.CriandoUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidades).Methods("Get")
	r.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8002", r))

}
