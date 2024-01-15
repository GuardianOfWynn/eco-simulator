package territory

type UpgradeLevel struct {
	Level uint8
	Cost  int32
	Value float32
}

type TerritoryUpgrade struct {
	Id          string
	Name        string
	Format      string
	Sprite      string
	UsedResorce ResourceType
	Levels      []BonusLevel
}

var (
	DEFENCE TerritoryUpgrade = TerritoryUpgrade{
		Id:          "defence",
		Name:        "Defence",
		Format:      "%",
		Sprite:      "",
		UsedResorce: FISH,
		Levels: []BonusLevel{
			{Level: 0, Cost: 0, Value: 0},
			{Level: 1, Cost: 100, Value: 300},
			{Level: 2, Cost: 300, Value: 450},
			{Level: 3, Cost: 600, Value: 525},
			{Level: 4, Cost: 1200, Value: 600},
			{Level: 5, Cost: 2400, Value: 650},
			{Level: 6, Cost: 4800, Value: 690},
			{Level: 7, Cost: 8400, Value: 720},
			{Level: 8, Cost: 12000, Value: 740},
			{Level: 9, Cost: 15600, Value: 760},
			{Level: 10, Cost: 19200, Value: 780},
			{Level: 11, Cost: 22800, Value: 800},
		},
	}

	DAMAGE TerritoryUpgrade = TerritoryUpgrade{
		Id:          "damage",
		Name:        "Damage",
		Format:      "%",
		Sprite:      "",
		UsedResorce: ORE,
		Levels: []BonusLevel{
			{Level: 0, Cost: 0, Value: 0},
			{Level: 1, Cost: 100, Value: 40},
			{Level: 2, Cost: 300, Value: 80},
			{Level: 3, Cost: 600, Value: 120},
			{Level: 4, Cost: 1200, Value: 160},
			{Level: 5, Cost: 2400, Value: 200},
			{Level: 6, Cost: 4800, Value: 240},
			{Level: 7, Cost: 8400, Value: 280},
			{Level: 8, Cost: 12000, Value: 320},
			{Level: 9, Cost: 15600, Value: 360},
			{Level: 10, Cost: 19200, Value: 400},
			{Level: 11, Cost: 22800, Value: 440},
		},
	}

	ATTACK_SPEED TerritoryUpgrade = TerritoryUpgrade{
		Id:          "attack_speed",
		Name:        "Attack Speed",
		Format:      "%",
		Sprite:      "",
		UsedResorce: CROP,
		Levels: []BonusLevel{
			{Level: 0, Cost: 0, Value: 0},
			{Level: 1, Cost: 100, Value: 50},
			{Level: 2, Cost: 300, Value: 100},
			{Level: 3, Cost: 600, Value: 150},
			{Level: 4, Cost: 1200, Value: 220},
			{Level: 5, Cost: 2400, Value: 300},
			{Level: 6, Cost: 4800, Value: 400},
			{Level: 7, Cost: 8400, Value: 500},
			{Level: 8, Cost: 12000, Value: 620},
			{Level: 9, Cost: 15600, Value: 660},
			{Level: 10, Cost: 19200, Value: 740},
			{Level: 11, Cost: 22800, Value: 840},
		},
	}

	HEALTH TerritoryUpgrade = TerritoryUpgrade{
		Id:          "health",
		Name:        "Health",
		Format:      "%",
		Sprite:      "",
		UsedResorce: WOOD,
		Levels: []BonusLevel{
			{Level: 0, Cost: 0, Value: 0},
			{Level: 1, Cost: 100, Value: 50},
			{Level: 2, Cost: 300, Value: 100},
			{Level: 3, Cost: 600, Value: 150},
			{Level: 4, Cost: 1200, Value: 220},
			{Level: 5, Cost: 2400, Value: 300},
			{Level: 6, Cost: 4800, Value: 400},
			{Level: 7, Cost: 8400, Value: 520},
			{Level: 8, Cost: 12000, Value: 640},
			{Level: 9, Cost: 15600, Value: 760},
			{Level: 10, Cost: 19200, Value: 880},
			{Level: 11, Cost: 22800, Value: 1000},
		},
	}
)

var Upgrades = map[string]TerritoryUpgrade{
	"attack_speed": ATTACK_SPEED,
	"defence":      DEFENCE,
	"damage":       DAMAGE,
	"health":       HEALTH,
}
