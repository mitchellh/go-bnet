package bnet

import(
	"fmt"
	"net/http"
	"testing"
	"reflect"
	"github.com/nmccrory/go-bnet/sc2"
	"github.com/nmccrory/go-bnet/wow"
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
const wowCharactersResp = `{ "characters":
				[{
				    "name": "foobar",
				    "realm": "foobar",
				    "battleGroup": "Foo",
				    "class": 1,
				    "race": 1,
				    "gender": 0,
				    "level": 99,
				    "achievementPoints": 1234,
				    "thumbnail": "foobar/123/avatar.jpg",
				    "spec": {
					"name": "foobar",
					"role": "foobar",
					"backgroundImage": "foo-bar",
					"icon": "foo_bar",
					"description": "Quick brown fox jumps over the lazy dog.",
					"order": 1
				    },
				    "guild": "Foo",
				    "guildRealm": "foobar",
				    "lastModified": 1234567
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
		t.Fatal("err: This user has no Starcraft 2 profile.")
	}
	want := sc2.SC2Character{ID: 1234567,
		Realm: 1,
		DisplayName: "foobar",
		ClanName: "foobar",
		ClanTag: "foobar",
		ProfilePath: "/profile/1234567/1/foobar/",
		Portrait: sc2.CharacterImage{-10, -10, 10, 10, 10,
			"http://media.blizzard.com/sc2/portraits/dummy.jpg"},
		Career: sc2.Career{"PROTOSS", 0, 0, 0,
			"DIAMOND", 0, 100},
		SwarmLevels: sc2.SwarmLevels{10,
			sc2.Level{1, 1000, 0},
			sc2.Level{2, 1000, 0},
			sc2.Level{3, 1000, 0}},
		Season: sc2.Season{123, 1, 2017, 0},
		Rewards: sc2.Rewards{[]int{12345678, 12345678}, []int{12345678, 12345678}},
		Achievements: sc2.Achievements{sc2.Points{1234},
			[]sc2.Achievement{sc2.Achievement{123456789, 123456789}}},
	}
	if !reflect.DeepEqual(actual.Characters[0], want) {
		t.Fatalf("returned %+v, want %+v", actual.Characters[0], want)
	}
}

func TestProfileService_WoW(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/wow/user/characters", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, wowCharactersResp)
	})
	actual, _, err := client.Profile().WoW()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if actual.Characters == nil {
		t.Fatal("err: This user has no World of Warcraft characters.")
	}
	want := wow.WoWCharacter{
		Name: "foobar",
		Realm: "foobar",
		BattleGroup: "Foo",
		Class: 1,
		Race: 1,
		Gender: 0,
		Level: 99,
		AchievementPoints: 1234,
		Thumbnail: "foobar/123/avatar.jpg",
		Spec: wow.Spec{"foobar", "foobar", "foo-bar", "foo_bar",
		"Quick brown fox jumps over the lazy dog.", 1},
		Guild: "Foo",
		GuildRealm: "foobar",
		LastModified: 1234567 }
	if !reflect.DeepEqual(actual.Characters[0], want) {
		t.Fatalf("returned %+v, want %+v", actual.Characters[0], want)
	}
}
