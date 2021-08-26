package api

import (
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-tip-api/internal/repo"
	desc "github.com/ozoncp/ocp-tip-api/pkg/ocp-tip-api"
)

type api struct {
	r repo.Repo
	p sarama.SyncProducer
	desc.UnimplementedOcpTipApiServer
}

func NewOcpTipApi(r repo.Repo, p sarama.SyncProducer) desc.OcpTipApiServer {
	return &api{r: r, p: p}
}
