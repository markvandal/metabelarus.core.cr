package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers mbpasstrust-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
		r.HandleFunc("/mbpasstrust/allowance", createAllowanceHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/mbpasstrust/allowance", listAllowanceHandler(cliCtx, "mbpasstrust")).Methods("GET")
		r.HandleFunc("/mbpasstrust/allowance/{key}", getAllowanceHandler(cliCtx, "mbpasstrust")).Methods("GET")
		r.HandleFunc("/mbpasstrust/allowance", setAllowanceHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/mbpasstrust/allowance", deleteAllowanceHandler(cliCtx)).Methods("DELETE")

		
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
