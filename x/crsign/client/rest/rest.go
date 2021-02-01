package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers crsign-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/crsign/auths/{id}", getAuthHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/crsign/auths", listAuthHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/crsign/id2services/{id}", getId2ServiceHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/crsign/records/{id}", getRecordHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/crsign/auth", requestAuthHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/crsign/auth/confirm", confirmAuthHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/crsign/records", createRecordHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/crsign/records/{id}", updateRecordHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/crsign/records/{id}", deleteRecordHandler(clientCtx)).Methods("POST")
}
