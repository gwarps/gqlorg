package types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/touchps/hackernews/config"
)

// Note that omitempty in name is important, which means we can use department object with Team
// without worrying of override and send it with team,
// same goes for other attributes
type Department struct {
	UID   string   `json:"uid,omitempty"`
	Name  string   `json:"name,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

func (deparment *Department) Create() (string, error) {
	deparment.DType = []string{"zion.department"}

	data, err := json.Marshal(deparment)
	if err != nil {
		return "", err
	}

	// op := &api.Operation{}
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
