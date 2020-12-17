package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers mbgovperm-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
		r.HandleFunc("/mbgovperm/consent", createConsentHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/mbgovperm/consent", listConsentHandler(cliCtx, "mbgovperm")).Methods("GET")
		r.HandleFunc("/mbgovperm/consent/{key}", getConsentHandler(cliCtx, "mbgovperm")).Methods("GET")
		r.HandleFunc("/mbgovperm/consent", setConsentHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/mbgovperm/consent", deleteConsentHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/mbgovperm/extservice", createExtserviceHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/mbgovperm/extservice", listExtserviceHandler(cliCtx, "mbgovperm")).Methods("GET")
		r.HandleFunc("/mbgovperm/extservice/{key}", getExtserviceHandler(cliCtx, "mbgovperm")).Methods("GET")
		r.HandleFunc("/mbgovperm/extservice", setExtserviceHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/mbgovperm/extservice", deleteExtserviceHandler(cliCtx)).Methods("DELETE")

		
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
