package types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/touchps/hackernews/config"
)

type Service struct {
	UID         string   `json:"uid,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Team        Team     `json:"zion.service.team,omitempty"`
	DType       []string `json:"dgraph.type,omitempty"`
}

func (service *Service) Create() (string, error) {
	service.DType = []string{"zion.service"}

	// Ensure only department ID is set nothing else
	data, err := json.Marshal(service)
	if err != nil {
		return "", err
	}
	ctx := context.Background()

	mu := &api.Mutation{
		CommitNow: true,
	}

	mu.SetJson = data

	response, err := config.DgraphClient.NewTxn().Mutate(ctx, mu)
	if err != nil {
		return "", err
	}

	fmt.Printf("%+v\n", response.Uids)
	var uids []string

	for _, val := range response.Uids {
		uids = append(uids, val)
	}

	if len(uids) != 1 {
		return "", errors.New(fmt.Sprintf("Created records not consistent %d", len(uids)))
	}

	return uids[0], nil
}
