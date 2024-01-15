package territory

type BonusLevel struct {
	Level          uint8
	Cost           int32
	Value          float32
	MaxTerritories int8
}

type TerritoryBonus struct {
	Id             string
	Name           string
	Format         string
	Sprite         string
	UsedResorce    ResourceType
	Levels         map[uint8]BonusLevel
	MaxTerritories uint8
}

var (
	STRONGER_MINIONS TerritoryBonus = TerritoryBonus{
		Id:          "stronger_minions",
		Name:        "Stronger Minions",
		Format:      "%",
		Sprite:      "",
		UsedResorce: WOOD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 200, Value: 150},
			2: {Level: 2, Cost: 400, Value: 200},
			3: {Level: 3, Cost: 800, Value: 250},
			4: {Level: 4, Cost: 1600, Value: 300},
		},
	}

	MULTI_HIT TerritoryBonus = TerritoryBonus{
		Id:             "multihit",
		Name:           "Tower Multi-Attacks",
		Format:         "Targets",
		Sprite:         "",
		MaxTerritories: 5,
		UsedResorce:    FISH,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 1},
			1: {Level: 1, Cost: 4800, Value: 2},
		},
	}

	TOWER_AURA TerritoryBonus = TerritoryBonus{
		Id:          "tower_aura",
		Name:        "Tower Aura",
		Format:      "s",
		Sprite:      "",
		UsedResorce: CROP,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 800, Value: 24},
			2: {Level: 2, Cost: 1600, Value: 18},
			3: {Level: 3, Cost: 3200, Value: 12},
		},
	}

	TOWER_VOLLEY TerritoryBonus = TerritoryBonus{
		Id:          "tower_volley",
		Name:        "Tower Volley",
		Format:      "s",
		Sprite:      "",
		UsedResorce: ORE,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 200, Value: 20},
			2: {Level: 2, Cost: 400, Value: 15},
			3: {Level: 3, Cost: 800, Value: 10},
		},
	}

	GATHERING_XP TerritoryBonus = TerritoryBonus{
		Id:          "gather_xp",
		Name:        "Gathering Experience",
		Format:      "%",
		Sprite:      "",
		UsedResorce: WOOD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 600, Value: 10},
			2: {Level: 2, Cost: 1300, Value: 20},
			3: {Level: 3, Cost: 2000, Value: 30},
			4: {Level: 4, Cost: 2700, Value: 40},
			5: {Level: 5, Cost: 3400, Value: 50},
			6: {Level: 6, Cost: 5500, Value: 60},
			7: {Level: 7, Cost: 10000, Value: 80},
			8: {Level: 8, Cost: 20000, Value: 100},
		},
	}

	MOB_XP TerritoryBonus = TerritoryBonus{
		Id:             "mob_xp",
		Name:           "Mob Experience",
		Format:         "%",
		Sprite:         "",
		MaxTerritories: 5,
		UsedResorce:    FISH,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 600, Value: 10},
			2: {Level: 2, Cost: 1200, Value: 20},
			3: {Level: 3, Cost: 1800, Value: 30},
			4: {Level: 4, Cost: 2400, Value: 40},
			5: {Level: 5, Cost: 3000, Value: 50},
			6: {Level: 6, Cost: 5000, Value: 60},
			7: {Level: 7, Cost: 10000, Value: 80},
			8: {Level: 8, Cost: 20000, Value: 100},
		},
	}

	MOB_DAMAGE TerritoryBonus = TerritoryBonus{
		Id:          "mob_damage",
		Name:        "Mob Damage",
		Format:      "%",
		Sprite:      "",
		UsedResorce: CROP,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 600, Value: 10},
			2: {Level: 2, Cost: 1200, Value: 20},
			3: {Level: 3, Cost: 1800, Value: 40},
			4: {Level: 4, Cost: 2400, Value: 60},
			5: {Level: 5, Cost: 3000, Value: 80},
			6: {Level: 6, Cost: 5000, Value: 120},
			7: {Level: 7, Cost: 10000, Value: 160},
			8: {Level: 8, Cost: 20000, Value: 200},
		},
	}

	PVP_DAMAGE TerritoryBonus = TerritoryBonus{
		Id:          "pvp_damage",
		Name:        "PvP Damage",
		Format:      "%",
		Sprite:      "",
		UsedResorce: ORE,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 600, Value: 5},
			2: {Level: 2, Cost: 1200, Value: 10},
			3: {Level: 3, Cost: 1800, Value: 15},
			4: {Level: 4, Cost: 2400, Value: 20},
			5: {Level: 5, Cost: 3000, Value: 25},
			6: {Level: 6, Cost: 5000, Value: 40},
			7: {Level: 7, Cost: 10000, Value: 60},
			8: {Level: 8, Cost: 20000, Value: 80},
		},
	}

	XP_SEEKING TerritoryBonus = TerritoryBonus{
		Id:          "xp_seeking",
		Name:        "XP Seeking",
		Format:      "/h",
		Sprite:      "",
		UsedResorce: EMERALD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 100, Value: 36000},
			2: {Level: 2, Cost: 200, Value: 66000},
			3: {Level: 3, Cost: 400, Value: 120000},
			4: {Level: 4, Cost: 800, Value: 228000},
			5: {Level: 5, Cost: 1600, Value: 456000},
			6: {Level: 6, Cost: 3200, Value: 900000},
			7: {Level: 7, Cost: 6400, Value: 1740000},
			8: {Level: 8, Cost: 9600, Value: 2580000},
			9: {Level: 9, Cost: 12800, Value: 3360000},
		},
	}

	TOME_SEEKING TerritoryBonus = TerritoryBonus{
		Id:          "tome_seeking",
		Name:        "Tome Seeking",
		Format:      "%/h",
		Sprite:      "",
		UsedResorce: FISH,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 800, Value: 0.15},
			2: {Level: 2, Cost: 4800, Value: 1.2},
			3: {Level: 3, Cost: 12800, Value: 2.4},
		},
	}

	EMERALD_SEEKING TerritoryBonus = TerritoryBonus{
		Id:          "emerald_seeking",
		Name:        "Emerald Seeking",
		Format:      "%/h",
		Sprite:      "",
		UsedResorce: WOOD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 200, Value: 0.3},
			2: {Level: 2, Cost: 800, Value: 3},
			3: {Level: 3, Cost: 1600, Value: 6},
			4: {Level: 4, Cost: 3200, Value: 12},
			5: {Level: 5, Cost: 6400, Value: 24},
		},
	}

	LARGER_RESOURCE_STORAGE TerritoryBonus = TerritoryBonus{
		Id:          "larger_resource_storage",
		Name:        "Larger Resource Storage",
		Format:      "%",
		Sprite:      "",
		UsedResorce: EMERALD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 400, Value: 100},
			2: {Level: 2, Cost: 800, Value: 300},
			3: {Level: 3, Cost: 2000, Value: 700},
			4: {Level: 4, Cost: 5000, Value: 1400},
			5: {Level: 5, Cost: 16000, Value: 3300},
			6: {Level: 6, Cost: 48000, Value: 7900},
		},
	}

	LARGER_EMERALD_STORAGE TerritoryBonus = TerritoryBonus{
		Id:          "larger_emerald_storage",
		Name:        "Larger Emerald Storage",
		Format:      "%",
		Sprite:      "",
		UsedResorce: WOOD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 200, Value: 100},
			2: {Level: 2, Cost: 400, Value: 300},
			3: {Level: 3, Cost: 1000, Value: 700},
			4: {Level: 4, Cost: 2500, Value: 1400},
			5: {Level: 5, Cost: 8000, Value: 3300},
			6: {Level: 6, Cost: 24000, Value: 7900},
		},
	}

	EFFICIENT_RESOURCES TerritoryBonus = TerritoryBonus{
		Id:          "efficient_resource",
		Name:        "Efficient Resources",
		Format:      "%",
		Sprite:      "",
		UsedResorce: EMERALD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 6000, Value: 50},
			2: {Level: 2, Cost: 12000, Value: 100},
			3: {Level: 3, Cost: 24000, Value: 150},
			4: {Level: 4, Cost: 48000, Value: 200},
			5: {Level: 5, Cost: 96000, Value: 250},
			6: {Level: 6, Cost: 192000, Value: 300},
		},
	}

	EFFICIENT_EMERALDS TerritoryBonus = TerritoryBonus{
		Id:          "efficient_emerald",
		Name:        "Efficient Emeralds",
		Format:      "%",
		Sprite:      "",
		UsedResorce: ORE,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 0},
			1: {Level: 1, Cost: 2000, Value: 35},
			2: {Level: 2, Cost: 8000, Value: 100},
			3: {Level: 3, Cost: 32000, Value: 300},
		},
	}

	RESOURCE_RATE TerritoryBonus = TerritoryBonus{
		Id:          "resource_rate",
		Name:        "Resource Rate",
		Format:      "s",
		Sprite:      "",
		UsedResorce: EMERALD,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 4},
			1: {Level: 1, Cost: 6000, Value: 3},
			2: {Level: 2, Cost: 18000, Value: 2},
			3: {Level: 3, Cost: 32000, Value: 1},
		},
	}

	EMERALD_RATE TerritoryBonus = TerritoryBonus{
		Id:          "emerald_rate",
		Name:        "Emerald Rate",
		Format:      "s",
		Sprite:      "",
		UsedResorce: CROP,
		Levels: map[uint8]BonusLevel{
			0: {Level: 0, Cost: 0, Value: 4},
			1: {Level: 1, Cost: 2000, Value: 3},
			2: {Level: 2, Cost: 8000, Value: 2},
			3: {Level: 3, Cost: 32000, Value: 1},
		},
	}
)

const (
	KEY_BONUS_STRONGER_MINIONS       = "stronger_minions"
	KEY_BONUS_MULTIHIT               = "multihit"
	KEY_BONUS_TOWER_AURA             = "tower_aura"
	KEY_BONUS_TOWER_VOLLEY           = "tower_volley"
	KEY_BONUS_GATHER_XP              = "gather_xp"
	KEY_BONUS_MOB_XP                 = "mob_xp"
	KEY_BONUS_MOB_DAMAGE             = "mob_damage"
	KEY_BONUS_PVP_DAMAGE             = "pvp_damage"
	KEY_BONUS_XP_SEEKING             = "xp_seeking"
	KEY_BONUS_TOME_SEEKING           = "tome_seeking"
	KEY_BONUS_EMERALD_SEEKING        = "emerald_seeking"
	KEY_BONUS_LARGE_RESOURCE_STORAGE = "large_resource_storage"
	KEY_BONUS_LARGE_EMERALD_STORAGE  = "larger_emerald_storage"
	KEY_BONUS_EFFICIENT_RESOURCES    = "efficient_resource"
	KEY_BONUS_EFFICIENT_EMERALDS     = "efficient_emerald"
	KEY_BONUS_RESOURCE_RATE          = "resource_rate"
	KEY_BONUS_EMERALD_RATE           = "emerald_rate"
)

var Bonuses = map[string]TerritoryBonus{
	"stronger_minions":        STRONGER_MINIONS,
	"multihit":                MULTI_HIT,
	"tower_aura":              TOWER_AURA,
	"tower_volley":            TOWER_VOLLEY,
	"gather_xp":               GATHERING_XP,
	"mob_xp":                  MOB_XP,
	"mob_damage":              MOB_DAMAGE,
	"pvp_damage":              PVP_DAMAGE,
	"xp_seeking":              XP_SEEKING,
	"tome_seeking":            TOME_SEEKING,
	"emerald_seeking":         EMERALD_SEEKING,
	"larger_resource_storage": LARGER_RESOURCE_STORAGE,
	"larger_emerald_storage":  LARGER_EMERALD_STORAGE,
	"efficient_resource":      EFFICIENT_RESOURCES,
	"efficient_emerald":       EFFICIENT_EMERALDS,
	"resource_rate":           RESOURCE_RATE,
	"emerald_rate":            EMERALD_RATE,
}
