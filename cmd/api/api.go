package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/humblgod/belajar-golang-rest-api/services/users"
)

type APIServer struct {
	addr string
	db	*sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db : db,
	}
}

func (s *APIServer) Run() error {
	// create router
	router := mux.NewRouter()

	// register servies
	userStore := users.NewUserStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegistersRoutes(router)



	//? config cors (*for fullstack)
	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"http://localhost:8080"}, 	// depends frontend router
	// 	AllowedMethods: []string{"GET", "POST", "DELETE"},
	// 	AllowedHeaders: []string{"Authorization", "Content-type"},
	// 	AllowCredentials: true,
	// })

	// corsHandler := c.Handler(router)

	// !log
	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, router)
}