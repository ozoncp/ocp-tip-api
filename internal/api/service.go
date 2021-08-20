package api

import (
	"github.com/ozoncp/ocp-tip-api/internal/repo"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
)

type api struct {
	r repo.Repo
	desc.UnimplementedOcpTipApiServer
}

func NewOcpTipApi(r repo.Repo) desc.OcpTipApiServer {
	return &api{r: r}
}
