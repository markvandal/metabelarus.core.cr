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

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
    r.HandleFunc("/crsign/id2auths/{id}", getId2AuthHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/crsign/id2auths", listId2AuthHandler(clientCtx)).Methods("GET")

    r.HandleFunc("/crsign/auths/{id}", getAuthHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/crsign/auths", listAuthHandler(clientCtx)).Methods("GET")

    r.HandleFunc("/crsign/id2signs/{id}", getId2SignHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/crsign/id2signs", listId2SignHandler(clientCtx)).Methods("GET")

    r.HandleFunc("/crsign/signatures/{id}", getSignatureHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/crsign/signatures", listSignatureHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
    r.HandleFunc("/crsign/id2auths", createId2AuthHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/id2auths/{id}", updateId2AuthHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/id2auths/{id}", deleteId2AuthHandler(clientCtx)).Methods("POST")

    r.HandleFunc("/crsign/auths", createAuthHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/auths/{id}", updateAuthHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/auths/{id}", deleteAuthHandler(clientCtx)).Methods("POST")

    r.HandleFunc("/crsign/id2signs", createId2SignHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/id2signs/{id}", updateId2SignHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/id2signs/{id}", deleteId2SignHandler(clientCtx)).Methods("POST")

    r.HandleFunc("/crsign/signatures", createSignatureHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/signatures/{id}", updateSignatureHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/crsign/signatures/{id}", deleteSignatureHandler(clientCtx)).Methods("POST")

}

