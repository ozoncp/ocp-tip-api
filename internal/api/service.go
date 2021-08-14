package api

import desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"

type api struct {
	desc.UnimplementedOcpTipApiServer
}

func NewOcpTipApi() desc.OcpTipApiServer {
	return &api{}
}
