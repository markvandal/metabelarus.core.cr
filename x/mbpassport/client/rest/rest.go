package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers mbpassport-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
		r.HandleFunc("/mbpassport/record", createRecordHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/mbpassport/record", listRecordHandler(cliCtx, "mbpassport")).Methods("GET")
		r.HandleFunc("/mbpassport/record/{key}", getRecordHandler(cliCtx, "mbpassport")).Methods("GET")
		r.HandleFunc("/mbpassport/record", setRecordHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/mbpassport/record", deleteRecordHandler(cliCtx)).Methods("DELETE")

		
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
