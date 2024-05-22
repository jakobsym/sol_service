package main

import (
	"log"
	"net/http"

	"github.com/jakobsym/sol_service/token/internal/controller/token"
	httphandler "github.com/jakobsym/sol_service/token/internal/handler/http"
	"github.com/jakobsym/sol_service/token/internal/repository/memory"
)

func main() {
	log.Println("starting token data service")
	repo := memory.New()
	ctrl := token.New(repo)
	h := httphandler.New(ctrl)

	http.HandleFunc("/token", h.GetTokenDetails)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
