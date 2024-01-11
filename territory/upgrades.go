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
		Levels:      []BonusLevel{},
	}

	DAMAGE TerritoryUpgrade = TerritoryUpgrade{
		Id:          "damage",
		Name:        "Damage",
		Format:      "%",
		Sprite:      "",
		UsedResorce: ORE,
		Levels:      []BonusLevel{},
	}

	ATTACK_SPEED TerritoryUpgrade = TerritoryUpgrade{
		Id:          "attack_speed",
		Name:        "Attack Speed",
		Format:      "%",
		Sprite:      "",
		UsedResorce: CROP,
		Levels:      []BonusLevel{},
	}

	HEALTH TerritoryUpgrade = TerritoryUpgrade{
		Id:          "health",
		Name:        "Health",
		Format:      "%",
		Sprite:      "",
		UsedResorce: WOOD,
		Levels:      []BonusLevel{},
	}
)

var Upgrades = map[string]TerritoryUpgrade{
	"attack_speed": ATTACK_SPEED,
	"defence":      DEFENCE,
	"damage":       DAMAGE,
	"health":       HEALTH,
}
