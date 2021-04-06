package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers crconsent-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/crconsent/requests/{id}", getRequestHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/crconsent/requests", listRequestHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/crconsent/requests", createRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/crconsent/requests/{id}", updateRequestHandler(clientCtx)).Methods("POST")
	// r.HandleFunc("/crconsent/requests/{id}", deleteRequestHandler(clientCtx)).Methods("POST")

}
