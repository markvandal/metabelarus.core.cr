package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    // this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers mbcorecr-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("/mbcorecr/identities/{id}", getIdentityHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/mbcorecr/identities", listIdentityHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/mbcorecr/identities", createIdentityHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/mbcorecr/identities/{id}", updateIdentityHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/mbcorecr/identities/{id}", deleteIdentityHandler(clientCtx)).Methods("POST")

}

