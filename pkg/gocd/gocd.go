package gocd

import (
	"context"
	"github.com/riksa/gonkka/pkg/gocd/client"
	"github.com/riksa/gonkka/pkg/gocd/client/operations"
	"github.com/riksa/gonkka/pkg/gocd/models"
	log "github.com/sirupsen/logrus"
)

type Rest struct {
	client *client.Gocd
	ctx    context.Context
}

func NewRest(host string, ctx context.Context) *Rest {
	config := client.DefaultTransportConfig()
	config.Host = host

	rest := &Rest{
		client: client.NewHTTPClientWithConfig(nil, config),
		ctx: ctx,
	}

	return rest
}

func (r *Rest) History(app string, id int64) (*models.HistoryResponse, error) {
	params := operations.NewHistoryParamsWithContext(r.ctx)
	params.ID = id
	params.App = app

	ok, err := r.client.Operations.History(params)
	if err != nil {
		log.WithError(err).WithField("id", id).Error("history request error")
		return nil, err
	} else {
		bytes, _ := ok.Payload.MarshalBinary()
		j := string(bytes)
		log.WithField("id", id).WithField("response", j).Debug("history request success")
		return ok.Payload, err
	}

}

func (r *Rest) Instance(app string, id int64) (*models.InstanceResponse, error) {
	params := operations.NewInstanceParamsWithContext(r.ctx)
	params.ID = id
	params.App = app

	ok, err := r.client.Operations.Instance(params)
	if err != nil {
		log.WithError(err).WithField("id", id).Error("instance request error")
		return nil, err
	} else {
		bytes, _ := ok.Payload.MarshalBinary()
		j := string(bytes)
		log.WithField("id", id).WithField("response", j).Debug("instance request success")
		return ok.Payload, err
	}

}