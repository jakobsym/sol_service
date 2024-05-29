package main

import (
	"log"
	"net/http"

	"github.com/jakobsym/sol_service/sol_service/internal/controller/sol_service"
	walletgateway "github.com/jakobsym/sol_service/sol_service/internal/gateway/wallet/http"
	httphandler "github.com/jakobsym/sol_service/sol_service/internal/handler/http"
)

func main() {
	log.Println("solservice service running at port 8083")
	walletGw := walletgateway.New("localhost:8082")
	ctrl := sol_service.New(walletGw)
	h := httphandler.New(ctrl)
	http.Handle("/solservice", http.HandlerFunc(h.GetAccountDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
