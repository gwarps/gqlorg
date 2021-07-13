package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/touchps/hackernews/graph/generated"
	"github.com/touchps/hackernews/graph/model"
	"github.com/touchps/hackernews/types"
)

func (r *departmentResolver) Teams(ctx context.Context, obj *model.Department) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, input model.NewDepartment) (*model.Department, error) {
	// panic(fmt.Errorf("not implemented"))

	var department model.Department
	typeDepartment := &types.Department{
		Name: input.Name,
	}
	id, err := typeDepartment.Create()

	if err != nil {
		return nil, err
	}

	department.Name = input.Name
	department.ID = id

	return &department, nil
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*model.Team, error) {
	// panic(fmt.Errorf("not implemented"))
	var team model.Team
	typeTeam := &types.Team{
		Name: input.Name,
		Department: types.Department{
			UID: input.DeparmentID,
		},
	}

	id, err := typeTeam.Create()
	if err != nil {
		return nil, err
	}

	team.Name = input.Name
	team.ID = id
	team.Deparment = &model.Department{
		ID: input.DeparmentID,
	}

	return &team, nil
}

func (r *mutationResolver) CreateService(ctx context.Context, input *model.NewService) (*model.Service, error) {
	var service model.Service
	typeService := &types.Service{
		Name:        input.Name,
		Description: input.Description,
		Team: types.Team{
			UID: input.TeamID,
		},
	}

	id, err := typeService.Create()
	if err != nil {
		return nil, err
	}
	service.Name = input.Name
	service.Description = input.Description
	service.ID = id

	return &service, nil
}

func (r *queryResolver) Deparments(ctx context.Context) ([]*model.Department, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Services(ctx context.Context) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *teamResolver) Services(ctx context.Context, obj *model.Team) ([]*model.Service, error) {
	panic(fmt.Errorf("not implemented"))
}

// Department returns generated.DepartmentResolver implementation.
func (r *Resolver) Department() generated.DepartmentResolver { return &departmentResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type departmentResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type teamResolver struct{ *Resolver }
