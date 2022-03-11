package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/internal/server/graph/generated"
	"github.com/cod3rcarl/wwdatabase-go-backend/graphql/internal/server/graph/model"
)

func (r *mutationResolver) CreateChampion(ctx context.Context, input model.CreateChampionInput) (*model.CreateChampionPayload, error) {
	pbChamp, err := r.Server.CreateChampion(ctx, &input)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating champ[ion")
	}

	return pbChamp, nil
}

func (r *mutationResolver) UpdateChampion(ctx context.Context, input model.UpdateChampionInput) (*model.UpdateChampionPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteChampion(ctx context.Context, input model.DeleteChampionInput) (*model.DeleteChampionPayload, error) {
	pbChamp, err := r.Server.DeleteChampion(ctx, &input)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating champ[ion")
	}

	return pbChamp, nil
}

func (r *queryResolver) Champion(ctx context.Context, titleHolder *string, currentChampion *bool, dateFilter *string, previousChampion *int) (*model.Champion, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Champions(ctx context.Context, filter *string, orderBy *model.ChampionOrderByInput) (*model.ChampionPayload, error) {
	pbChamp, err := r.Server.GetChampions(ctx, orderBy, filter)
	if err != nil {
		return nil, errors.Wrap(err, "Error creating champ[ion")
	}

	return pbChamp, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)
