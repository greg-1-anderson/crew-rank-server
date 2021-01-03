package graph

import "twta/crew-rank-server/graph/model"

//go:generate go run github.com/99designs/gqlgen

func playerAggrageteStats(name string, imposter *model.ImposterStats, crewmate *model.CrewmateStats) *model.PlayerAggragateStats {
	// Fill in / correct the total number of games
	imposter.Games.Total = imposter.Games.Won + imposter.Games.Lost
	crewmate.Games.Total = crewmate.Games.Won + crewmate.Games.Lost

	// Calculate the overall game statistics
	overall := model.AggragateStats{
		Won:   imposter.Games.Won + crewmate.Games.Won,
		Lost:  imposter.Games.Lost + crewmate.Games.Lost,
		Total: imposter.Games.Total + crewmate.Games.Total,
	}

	return &model.PlayerAggragateStats{
		Name: name,
		Overall: &model.OverallStats{
			Games: &overall,
		},
		Imposter: imposter,
		Crewmate: crewmate,
	}
}
