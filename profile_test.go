package bnet

import(
	"fmt"
	"net/http"
	"testing"
	"reflect"
)
const sc2ProfileResp = `{ "characters":
				[{
				    "id": 1234567,
				    "realm": 1,
				    "displayName": "foobar",
				    "clanName": "foobar",
				    "clanTag": "foobar",
				    "profilePath": "/profile/1234567/1/foobar/",
				    "portrait": {
					"x": -10,
					"y": -10,
					"w": 10,
					"h": 10,
					"offset": 10,
					"url": "http://media.blizzard.com/sc2/portraits/dummy.jpg"
				    },
				    "career": {
					"primaryRace": "PROTOSS",
					"terranWins": 0,
					"protossWins": 0,
					"zergWins": 0,
					"highest1v1Rank": "DIAMOND",
					"seasonTotalGames": 0,
					"careerTotalGames": 100
				    },
				    "swarmLevels": {
					"level": 10,
					"terran": {
					    "level": 1,
					    "totalLevelXP": 1000,
					    "currentLevelXP": 0
					},
					"zerg": {
					    "level": 2,
					    "totalLevelXP": 1000,
					    "currentLevelXP": 0
					},
					"protoss": {
					    "level": 3,
					    "totalLevelXP": 1000,
					    "currentLevelXP": 0
					}
				    },
				    "campaign": {},
				    "season": {
					"seasonId": 123,
					"seasonNumber": 1,
					"seasonYear": 2017,
					"totalGamesThisSeason": 0
				    },
				    "rewards": {
					"selected": [12345678, 12345678],
					"earned": [12345678, 12345678]
				    },
				    "achievements": {
					"points": {
					    "totalPoints": 1234,
					    "categoryPoints": {}
					},
					"achievements": [{
					    "achievementId": 123456789,
					    "completionDate": 123456789
					}]
				    }
				}]
			}`

func TestProfileService_SC2(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/sc2/profile/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, sc2ProfileResp)
	})
	actual, _, err := client.Profile().SC2()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual.Characters == nil {
		t.Fatal("err: Profile is empty -> &SC2Profile{Characters: []SC2Character{}(nil)}")
	}
	want := SC2Character{ID: 1234567,
		Realm: 1,
		DisplayName: "foobar",
		ClanName: "foobar",
		ClanTag: "foobar",
		ProfilePath: "/profile/1234567/1/foobar/",
		Portrait: CharacterImage{-10, -10, 10, 10, 10,
			"http://media.blizzard.com/sc2/portraits/dummy.jpg"},
		Career: Career{"PROTOSS", 0, 0, 0,
			"DIAMOND", 0, 100},
		SwarmLevels: SwarmLevels{10,
			Level{1, 1000, 0},
			Level{2, 1000, 0},
			Level{3, 1000, 0}},
		Season: Season{123, 1, 2017, 0},
		Rewards: Rewards{[]int{12345678, 12345678}, []int{12345678, 12345678}},
		Achievements: Achievements{Points{1234},
			[]Achievement{Achievement{123456789, 123456789}}},
	}
	if !reflect.DeepEqual(actual.Characters[0], want) {
		t.Fatalf("returned %+v, want %+v", actual.Characters[0], want)
	}
}
