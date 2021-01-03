package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"twta/crew-rank-server/graph/generated"
	"twta/crew-rank-server/graph/model"
)

func (r *queryResolver) CrewRank(ctx context.Context) (*model.CrewRank, error) {
	ladder := []*model.PlayerAggragateStats{}

	ladder = append(
		ladder,
		playerAggrageteStats(
			"The Dead Body",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  10,
					Lost: 50,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  40,
					Lost: 10,
				},
			},
		),
	)

	return &model.CrewRank{
		Ladder: ladder,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
