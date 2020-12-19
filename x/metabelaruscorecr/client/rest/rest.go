package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers metabelaruscorecr-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/metabelaruscorecr/confirmation", createConfirmationHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/metabelaruscorecr/confirmation", listConfirmationHandler(cliCtx, "metabelaruscorecr")).Methods("GET")
	r.HandleFunc("/metabelaruscorecr/confirmation/{key}", getConfirmationHandler(cliCtx, "metabelaruscorecr")).Methods("GET")
	r.HandleFunc("/metabelaruscorecr/confirmation", setConfirmationHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/metabelaruscorecr/confirmation", deleteConfirmationHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/metabelaruscorecr/invitation", createInvitationHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/metabelaruscorecr/invitation", listInvitationHandler(cliCtx, "metabelaruscorecr")).Methods("GET")
	r.HandleFunc("/metabelaruscorecr/invitation/{key}", getInvitationHandler(cliCtx, "metabelaruscorecr")).Methods("GET")
	r.HandleFunc("/metabelaruscorecr/invitation", setInvitationHandler(cliCtx)).Methods("PUT")
	r.HandleFunc("/metabelaruscorecr/invitation", deleteInvitationHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/metabelaruscorecr/Identity", createIdentityHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/metabelaruscorecr/Identity", listIdentityHandler(cliCtx, "metabelaruscorecr")).Methods("GET")
	r.HandleFunc("/metabelaruscorecr/Identity/{key}", getIdentityHandler(cliCtx, "metabelaruscorecr")).Methods("GET")
	r.HandleFunc("/metabelaruscorecr/Identity", setIdentityHandler(cliCtx)).Methods("PUT")
}
