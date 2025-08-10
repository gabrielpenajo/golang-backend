package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Users list..."))
}

func (a *api) createUsersHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Created user!"))
}

func main() {
	api := &api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	srv.ListenAndServe()
}
