package client

import (
	"github.com/spf13/cobra"

	"github.com/Dipper-Protocol/client/context"
	"github.com/Dipper-Protocol/codec"
	"github.com/Dipper-Protocol/x/gov/client/rest"
)

// function to create the rest handler
type RESTHandlerFn func(context.CLIContext) rest.ProposalRESTHandler

// function to create the cli handler
type CLIHandlerFn func(*codec.Codec) *cobra.Command

// The combined type for a proposal handler for both cli and rest
type ProposalHandler struct {
	CLIHandler  CLIHandlerFn
	RESTHandler RESTHandlerFn
}

// NewProposalHandler creates a new ProposalHandler object
func NewProposalHandler(cliHandler CLIHandlerFn, restHandler RESTHandlerFn) ProposalHandler {
	return ProposalHandler{
		CLIHandler:  cliHandler,
		RESTHandler: restHandler,
	}
}
