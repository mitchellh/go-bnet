package sc2

import (
	"github.com/mitchellh/go-bnet"
)

// TODO: Create a 'Campaign' struct to represent a character's campaign progress.

// CharacterImage is a character's portrait or avatar.
type CharacterImage struct {
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
	Offset int    `json:"offset"`
	Url    string `json:"url"`
}

// Career represents game statistics for a character's Battle.net career.
type Career struct {
	PrimaryRace	 string	`json:"primaryRace"`
	TerranWins	 int	`json:"terranWins"`
	ProtossWins	 int	`json:"protossWins"`
	ZergWins	 int	`json:"zergWins"`
	Highest1v1Rank	 string	`json:"highest1v1Rank"`
	SeasonTotalGames int	`json:"seasonTotalGames"`
	CareerTotalGames int	`json:"careerTotalGames"`
}

// Level is the current level and XP a character has earned.
type Level struct {
	Level	       int `json:"level"`
	TotalLevelXP   int `json:"totalLevelXP"`
	CurrentLevelXP int `json:"currentLevelXP"`
}

// SwarmLevels represents a character's level for each swarm (race) as well as their overall level.
type SwarmLevels struct {
	Level	int   `json:"level"`
	Terran	Level `json:"terran"`
	Zerg	Level `json:"zerg"`
	Protoss Level `json:"protoss"`
}

// Season is the current Starcraft 2 online multiplayer season.
type Season struct {
	ID 	   int `json:"seasonId"`
	Number     int `json:"seasonNumber"`
	Year       int `json:"seasonYear"`
	TotalGames int `json:"totalGamesThisSeason"`
}

// Rewards represents selected and earned rewards for a profile.
type Rewards struct {
	Selected []int `json:"selected"`
	Earned   []int `json:"earned"`
}

// Points holds a character's total achievement points.
type Points struct {
	Total int `json:"totalPoints"`
}

// Achievement represents a single Starcraft 2 achievement.
type Achievement struct {
	ID             int `json:"achievementId"`
	CompletionDate int `json:"completionDate"`
}

// Achievements represents achievement information for a Starcraft 2 profile.
type Achievements struct {
	Points       Points        `json:"points"`
	Achievements []Achievement `json:"achievements"`
}

// SC2Character represents a character in a user's Starcraft 2 profile.
type SC2Character struct {
	ID           int            `json:"id"`
	Realm        int            `json:"realm"`
	Name         string	    `json:"name"`
	DisplayName  string	    `json:"displayName"`
	ClanName     string	    `json:"clanName"`
	ClanTag      string	    `json:"clanTag"`
	ProfilePath  string	    `json:"profilePath"`
	Portrait     CharacterImage `json:"portrait"`
	Avatar	     CharacterImage `json:"avatar"`
	Career	     Career	    `json:"career"`
	SwarmLevels  SwarmLevels    `json:"swarmLevels"`
	Season	     Season         `json:"season"`
	Rewards      Rewards        `json:"rewards"`
	Achievements Achievements   `json:"achievements"`
}