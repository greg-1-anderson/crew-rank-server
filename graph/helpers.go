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

func mockLadder() []*model.PlayerAggragateStats {
	ladder := []*model.PlayerAggragateStats{
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
		playerAggrageteStats(
			"Innocent Crewmate",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  20,
					Lost: 30,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  50,
					Lost: 40,
				},
			},
		),
		playerAggrageteStats(
			"Pacifist",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  2,
					Lost: 3,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  5,
					Lost: 4,
				},
			},
		),
		playerAggrageteStats(
			"Vent Runner",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  13,
					Lost: 2,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  5,
					Lost: 18,
				},
			},
		),
		playerAggrageteStats(
			"Fast Talk",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  2,
					Lost: 19,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  21,
					Lost: 8,
				},
			},
		),
		playerAggrageteStats(
			"Minding My Own Business",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  0,
					Lost: 3,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  5,
					Lost: 1,
				},
			},
		),
		playerAggrageteStats(
			"Saboteur",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  25,
					Lost: 3,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  5,
					Lost: 15,
				},
			},
		),
		playerAggrageteStats(
			"Nobody Saw Me Do It",
			&model.ImposterStats{
				Games: &model.AggragateStats{
					Won:  60,
					Lost: 1,
				},
			},
			&model.CrewmateStats{
				Games: &model.AggragateStats{
					Won:  2,
					Lost: 19,
				},
			},
		),
	}

	return ladder
}
