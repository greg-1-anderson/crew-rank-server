package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"twta/crew-rank-server/graph/generated"
	"twta/crew-rank-server/graph/model"
)

func (r *queryResolver) CrewRank(ctx context.Context) (*model.CrewRank, error) {
	return &model.CrewRank{
		Ladder: mockLadder(),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
