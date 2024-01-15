package territory

import (
	"math"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RouteStyle uint8
type BorderStyle uint8
type ResourceType uint8
type Treasury uint8
type TransferDirection uint8
type Storage map[ResourceType]int64

var BaseTerritoriesMap = map[string]BaseTerritory{}

const (
	BASE_RESOURCE_PRODUCTION  = 1
	BASE_EMERALD_PRODUCTION   = 2.5
	BASE_RESOURCE_STORAGE     = 300
	BASE_EMERALD_STORAGE      = 3000
	BASE_HQ_DAMAGE            = 10000
	BASE_TERRITORY_DAMAGE     = 10000
	HQ_RESOURCE_STORAGE_BOOST = 5
	HQ_EMERALD_STORAGE_BOOST  = 1.666666
)

const (
	TERRITORY_TO_HQ TransferDirection = iota
	HQ_TO_TERRITORY
)

const (
	CHEAPEST RouteStyle = iota
	FASTEST
)

const (
	CROP ResourceType = iota
	ORE
	EMERALD
	WOOD
	FISH
)

const (
	CLOSED BorderStyle = iota
	OPEN
)

const (
	VERY_LOW Treasury = iota
	LOW
	MEDIUM
	HIGH
	VERY_HIGH
)

type ResourceTransference struct {
	Id        string
	Direction TransferDirection
	Storage   Storage
	Target    *Territory
}

type BaseTerritory struct {
	Name              string   `json:"territory"`
	OreMultiplier     float64  `json:"ore"`
	CropMultiplier    float64  `json:"crop"`
	WoodMultiplier    float64  `json:"wood"`
	FishMultiplier    float64  `json:"fish"`
	EmeraldMultiplier float64  `json:"emerald"`
	Conns             []string `json:"conns"`
}

func (b *BaseTerritory) CreateTerritoryInstance() *Territory {
	return &Territory{
		Name:                 b.Name,
		lastResourceProduced: 0,
		lastEmeraldProduced:  0,
		lastConsumedResource: 0,
		lastResourceTransfer: 0,
		resourceOverflow:     false,
		resourceGap:          false,
		HQ:                   false,
		Tax:                  0.05,
		AllyTax:              0.05,
		Treasury:             VERY_LOW,
		RouteStyle:           RouteStyle(OPEN),
		Claim:                nil,
		Borders:              CLOSED,
		TargetTerritory:      "",
		Connections:          b.Conns,
		Upgrades: map[string]uint8{
			"attack_speed": 1,
			"defence":      1,
			"damage":       1,
			"health":       1,
		},
		Storage: Storage{
			CROP:    0,
			EMERALD: 0,
			FISH:    0,
			ORE:     0,
			WOOD:    0,
		},
		PassingResource: []ResourceTransference{},
		ProductionMultipliers: map[ResourceType]float64{
			CROP:    b.CropMultiplier,
			EMERALD: b.EmeraldMultiplier,
			FISH:    b.FishMultiplier,
			WOOD:    b.WoodMultiplier,
			ORE:     b.OreMultiplier,
		},
		Bonuses: map[string]uint8{
			"stronger_minions":        0,
			"multihit":                0,
			"tower_aura":              0,
			"tower_volley":            0,
			"gather_xp":               0,
			"mob_xp":                  0,
			"mob_damage":              0,
			"pvp_damage":              0,
			"xp_seeking":              0,
			"tome_seeking":            0,
			"emerald_seeking":         0,
			"larger_resource_storage": 0,
			"larger_emerald_storage":  0,
			"efficient_resource":      0,
			"efficient_emerald":       0,
			"resource_rate":           0,
			"emerald_rate":            0,
		},
	}
}

type Territory struct {
	HQ                    bool
	Name                  string
	Claim                 *Claim
	Treasury              Treasury
	RouteStyle            RouteStyle
	Borders               BorderStyle
	Tax                   float64
	AllyTax               float64
	Storage               Storage
	ProductionMultipliers map[ResourceType]float64
	PassingResource       []ResourceTransference
	TargetTerritory       string
	lastResourceProduced  int64
	lastEmeraldProduced   int64
	lastConsumedResource  int64
	lastResourceTransfer  int64
	resourceOverflow      bool
	resourceGap           bool
	Bonuses               map[string]uint8
	Upgrades              map[string]uint8
	Connections           []string
}

func (t *Territory) Reset() {
	t.lastResourceProduced = 0
	t.lastEmeraldProduced = 0
	t.lastConsumedResource = 0
	t.lastResourceTransfer = 0
	t.resourceOverflow = false
	t.resourceGap = false
	t.HQ = false
	t.Treasury = VERY_LOW
	t.RouteStyle = RouteStyle(OPEN)
	t.Claim = nil
	t.Borders = CLOSED
	t.TargetTerritory = ""
	t.Upgrades = map[string]uint8{
		"attack_speed": 1,
		"defence":      1,
		"damage":       1,
		"health":       1,
	}
	t.Storage = Storage{
		CROP:    0,
		EMERALD: 0,
		FISH:    0,
		ORE:     0,
		WOOD:    0,
	}
	t.PassingResource = []ResourceTransference{}
	t.Bonuses = map[string]uint8{
		"stronger_minions":        0,
		"multihit":                0,
		"tower_aura":              0,
		"tower_volley":            0,
		"gather_xp":               0,
		"mob_xp":                  0,
		"mob_damage":              0,
		"pvp_damage":              0,
		"xp_seeking":              0,
		"tome_seeking":            0,
		"emerald_seeking":         0,
		"larger_resource_storage": 0,
		"larger_emerald_storage":  0,
		"efficient_resource":      0,
		"efficient_emerald":       0,
		"resource_rate":           0,
		"emerald_rate":            0,
	}
}

func (t *Territory) GetTowerDamageLow() int32 {
	return 0
}

func (t *Territory) GetTowerDamageHigh() int32 {
	return 0
}

func (t *Territory) GetTowerDefence() float32 {
	return 0
}

func (t *Territory) GetTowerAttackSpeed() float32 {
	return 0
}

func (t *Territory) GetTowerHP() int32 {
	return 0
}

func (t *Territory) GetResourceRate() int64 {
	return int64(Bonuses[KEY_BONUS_EMERALD_RATE].Levels[t.Bonuses[KEY_BONUS_RESOURCE_RATE]].Value)
}

func (t *Territory) GetEmeraldRate() int64 {
	return int64(Bonuses[KEY_BONUS_EMERALD_RATE].Levels[t.Bonuses[KEY_BONUS_EMERALD_RATE]].Value)
}

func (t *Territory) getTreasuryBonus() float64 {
	pfinder := Pathfinder{
		Root:     t,
		GuildMap: EngineInstance.Map,
	}
	treasuryBonus := 0.0
	if t.Claim != nil {
		distance := pfinder.GetDistance(t.Claim.GetHQ())
		baseBonus := float64(1 - math.Max(0, float64(distance)-2)*0.15)
		switch t.Treasury {
		case VERY_LOW:
			treasuryBonus = 0
			break
		case LOW:
			treasuryBonus = baseBonus
			break
		case MEDIUM:
			treasuryBonus = baseBonus * 2
			break
		case HIGH:
			treasuryBonus = baseBonus * 2.5
			break
		case VERY_HIGH:
			treasuryBonus = baseBonus * 3
			break
		}
	}
	return treasuryBonus
}

func (t *Territory) GetProducedEmerald() float64 {
	multiplier := int32(Bonuses[KEY_BONUS_EFFICIENT_EMERALDS].Levels[t.Bonuses[KEY_BONUS_EFFICIENT_EMERALDS]].Value)
	return (1 + float64(multiplier/100)) * BASE_EMERALD_PRODUCTION * (1 + t.getTreasuryBonus())
}

func (t *Territory) GetProducedResource() float64 {
	multiplier := int32(Bonuses[KEY_BONUS_EFFICIENT_RESOURCES].Levels[t.Bonuses[KEY_BONUS_EFFICIENT_RESOURCES]].Value) / 100
	return float64(1+multiplier) * BASE_RESOURCE_PRODUCTION * (1 + t.getTreasuryBonus())
}

func (t *Territory) GetEmeraldStorageSize() int32 {
	multiplier := int32(Bonuses[KEY_BONUS_LARGE_EMERALD_STORAGE].Levels[t.Bonuses[KEY_BONUS_LARGE_EMERALD_STORAGE]].Value)/100 + 1
	if t.HQ {
		return int32(math.Ceil(float64(float64(multiplier) * float64(BASE_EMERALD_STORAGE) * float64(HQ_EMERALD_STORAGE_BOOST))))
	}
	return (1 + multiplier) * BASE_EMERALD_STORAGE
}

func (t *Territory) GetResourceStorageSize() int32 {
	multiplier := int32(Bonuses[KEY_BONUS_LARGE_RESOURCE_STORAGE].Levels[t.Bonuses[KEY_BONUS_LARGE_RESOURCE_STORAGE]].Value)/100 + 1
	if t.HQ {
		return int32(math.Ceil(float64((multiplier) * BASE_RESOURCE_STORAGE * HQ_RESOURCE_STORAGE_BOOST)))
	}
	return (multiplier) * BASE_EMERALD_STORAGE
}

func (t *Territory) FindExternal() []*Territory {
	pf := Pathfinder{
		Root:     t,
		GuildMap: EngineInstance.Map,
	}
	externals := []*Territory{}
	for _, v := range EngineInstance.Map.Territories {
		if pf.GetDistance(v, FASTEST) <= 3 {
			externals = append(externals, v)
		}
	}
	return externals
}

// GetResourceCost: Retrieves the resource costs per second for territory
func (t *Territory) GetResourceCosts() Storage {
	costs := Storage{
		CROP:    0,
		EMERALD: 0,
		FISH:    0,
		ORE:     0,
		WOOD:    0,
	}
	for k, v := range t.Bonuses {
		bonus := Bonuses[k]
		level := bonus.Levels[v]
		costs[bonus.UsedResorce] = costs[bonus.UsedResorce] + int64(level.Cost)
	}
	return costs
}

func (t *Territory) Tick() {
	currentTimeMillis := time.Now().UnixMilli()
	if t.Claim != nil {
		// Produce emerald
		if currentTimeMillis-t.lastEmeraldProduced >= t.GetEmeraldRate()*1000 {
			t.lastEmeraldProduced = currentTimeMillis
			t.Storage[EMERALD] = t.Storage[EMERALD] + int64(t.GetProducedEmerald()*float64(t.ProductionMultipliers[EMERALD]))
		}

		// Produce resource
		if currentTimeMillis-t.lastResourceProduced >= t.GetResourceRate()*1000 {
			t.lastResourceProduced = currentTimeMillis
			for k, v := range t.Storage {
				if k == EMERALD {
					continue
				}
				t.Storage[k] = v + int64(t.GetProducedResource()*float64(t.ProductionMultipliers[k]))
			}
		}

		// Consume resources
		if currentTimeMillis-t.lastConsumedResource >= 1000 {
			t.lastConsumedResource = currentTimeMillis
			cost := t.GetResourceCosts()
			t.ConsumeResources(cost)
			for _, v := range cost {
				if v != 0 {
					t.Claim.AskForResources(t, cost)
					break
				}
			}
		}
	}
	// Transfer resource
	if currentTimeMillis-t.lastResourceTransfer >= 60000 {
		t.lastResourceTransfer = currentTimeMillis
		if t.Claim != nil && !t.HQ {
			transfer := ResourceTransference{
				Id:        uuid.NewString(),
				Target:    t.Claim.GetHQ(),
				Direction: TERRITORY_TO_HQ,
				Storage:   t.Storage,
			}
			t.TransferResource(transfer)
			t.ConsumeResources(t.Storage)
		}
		for _, r := range t.PassingResource {
			t.TransferResource(r)
		}
		t.PassingResource = []ResourceTransference{}
	}
}

func (t *Territory) TransferResource(transf ResourceTransference) {
	target := transf.Target
	if transf.Direction == TERRITORY_TO_HQ {
		if t.Claim != nil {
			// If territory is member of a claim, reroute it to HQ
			target = t.Claim.GetHQ()
			transf.Target = target
		}
	}
	t.TargetTerritory = transf.Target.Name
	for _, conn := range t.Connections {
		if strings.EqualFold(conn, target.Name) {
			target.ReceiveResource(transf)
			return
		}
	}
	if strings.EqualFold(t.Name, transf.Target.Name) {
		target.ReceiveResource(transf)
		return
	}
	pathfinder := Pathfinder{
		Root:     t,
		GuildMap: EngineInstance.Map,
	}
	route := pathfinder.Route(EngineInstance.Map.GetTerritory(t.TargetTerritory), t.RouteStyle)
	if len(route) > 0 {
		route[0].ReceiveResource(transf)
	}
}

func (t *Territory) ReceiveResource(transference ResourceTransference) {
	if t.Name == transference.Target.Name {
		t.StoreResource(transference.Storage)
	} else {
		t.PassingResource = append(t.PassingResource, transference)
	}
}

func (t *Territory) ConsumeResources(costs Storage) {
	for k, v := range costs {
		stored := t.Storage[k]
		if (stored - v) < 0 {
			for _, transference := range t.PassingResource {
				passing := transference.Storage[k]
				if passing+stored-v < 0 {
					t.resourceGap = true
				} else {
					t.Storage[k] = 0
					transference.Storage[k] = passing + stored - v
					t.resourceGap = false
				}
			}
		} else {
			t.resourceGap = false
			t.Storage[k] = t.Storage[k] - v
		}
	}
}

func (t *Territory) StoreResource(resources Storage) {
	storage := 0
	for k, v := range resources {
		if k == EMERALD {
			storage = int(t.GetEmeraldStorageSize())
		} else {
			storage = int(t.GetResourceStorageSize())
		}
		stored := t.Storage[k]
		if stored+v > int64(storage) {
			t.resourceOverflow = true
			t.Storage[k] = int64(storage)
		} else {
			t.resourceOverflow = false
			t.Storage[k] = stored + v
		}
	}
}
