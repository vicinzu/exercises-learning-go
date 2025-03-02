package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams []Team
	Wins  map[string]uint
}

func (l *League) MatchResult(firstTeamName string, firstTeamScore int, secondTeamName string, secondTeamScore int) {
	scoreDiff := firstTeamScore - secondTeamScore
	switch {
	case scoreDiff > 0:
		l.Wins[firstTeamName]++
	case scoreDiff < 0:
		l.Wins[secondTeamName]++
	}
}

func (l League) Ranking() []string {
	ranking := make([]string, 0, len(l.Teams))

	for _, team := range l.Teams {
		ranking = append(ranking, team.Name)
	}
	sort.Slice(ranking, func(i, j int) bool {
		return l.Wins[ranking[i]] > l.Wins[ranking[j]]
	})

	return ranking
}

func main() {
	league := League{
		Name: "Premier League",
		Teams: []Team{
			{
				Name:    "Manchester United",
				Players: []string{"Player 1", "Player 2", "Player 3"},
			},
			{
				Name:    "Manchester City",
				Players: []string{"Player 4", "Player 5", "Player 6"},
			},
			{
				Name:    "Chelsea",
				Players: []string{"Player 7", "Player 8", "Player 9"},
			},
			{
				Name:    "Liverpool",
				Players: []string{"Player 10", "Player 11", "Player 12"},
			},
			{
				Name:    "Arsenal",
				Players: []string{"Player 13", "Player 14", "Player 15"},
			},
		},
		Wins: map[string]uint{},
	}
	league.MatchResult("Manchester United", 2, "Manchester City", 1)
	league.MatchResult("Manchester United", 1, "Chelsea", 2)
	league.MatchResult("Manchester United", 3, "Liverpool", 0)
	league.MatchResult("Manchester United", 1, "Arsenal", 1)
	league.MatchResult("Manchester City", 2, "Chelsea", 2)
	league.MatchResult("Manchester City", 1, "Liverpool", 1)
	league.MatchResult("Manchester City", 3, "Arsenal", 0)
	league.MatchResult("Chelsea", 2, "Liverpool", 1)
	league.MatchResult("Chelsea", 1, "Arsenal", 1)
	league.MatchResult("Liverpool", 2, "Arsenal", 1)

	fmt.Printf("Ranking: %#v\n", league.Ranking())
}
