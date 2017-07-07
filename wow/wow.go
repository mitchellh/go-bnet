package wow

// Character represents a single World of Warcraft character.
type Character struct {
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	BattleGroup       string `json:"battlegroup"`
	Class             int    `json:"class"`
	Race              int    `json:"race"`
	Gender            int    `json:"gender"`
	Level             int    `json:"level"`
	AchievementPoints int    `json:"achievementPoints"`
	Thumbnail         string `json:"thumbnail"`
	Spec              Spec   `json:"spec"`
	Guild             string `json:"guild"`
	GuildRealm        string `json:"guildRealm"`
	LastModified      int    `json:"lastModified"`
}

// Spec represents a character's class spec
// (e.g. Frost Mage, Marksmanship Hunter, etc.)
type Spec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}



