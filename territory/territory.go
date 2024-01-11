package territory

import (
	"time"

	"github.com/victorbetoni/go-streams/streams"
)

type RouteStyle uint8
type BorderStyle uint8
type ResourceType uint8
type Treasury uint8

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

type Storage struct {
	StoredEmerald int32
	StoredOre     int32
	StoredCrop    int32
	StoredWood    int32
	StoredFish    int32
}

type BaseTerritory struct {
	Name              string   `json:"name"`
	OreMultiplier     float32  `json:"ore"`
	CropMultiplier    float32  `json:"crop"`
	WoodMultiplier    float32  `json:"wood"`
	FishMultiplier    float32  `json:"fish"`
	EmeraldMultiplier float32  `json:"emerald"`
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
		Treasury:             VERY_LOW,
		RouteStyle:           RouteStyle(OPEN),
		Borders:              CLOSED,
		TargetTerritory:      "",
		Upgrades: map[string]uint8{
			"attack_speed": 1,
			"defence":      1,
			"damage":       1,
			"health":       1,
		},
		Storage: map[ResourceType]int64{
			CROP:    0,
			EMERALD: 0,
			FISH:    0,
			ORE:     0,
			WOOD:    0,
		},
		PassingResource: map[ResourceType]int64{
			CROP:    0,
			EMERALD: 0,
			FISH:    0,
			ORE:     0,
			WOOD:    0,
		},
		ProductionMultipliers: map[ResourceType]float32{
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
	Storage               map[ResourceType]int64
	PassingResource       map[ResourceType]int64
	ProductionMultipliers map[ResourceType]float32
	TargetTerritory       string
	lastResourceProduced  int64
	lastEmeraldProduced   int64
	lastConsumedResource  int64
	lastResourceTransfer  int64
	resourceOverflow      bool
	resourceGap           bool
	Bonuses               map[string]uint8
	Upgrades              map[string]uint8
}

type PassingResource struct {
	Target    string
	Resources Storage
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
	return 0
}

func (t *Territory) GetEmeraldRate() int64 {
	return 0
}

func (t *Territory) GetProducedEmerald() int32 {
	return 0
}

func (t *Territory) GetProducedResource() int32 {
	return 0
}

func (t *Territory) GetEmeraldStorageSize() int32 {
	return 0
}

func (t *Territory) GetResourceStorageSize() int32 {
	return 0
}

// GetResourceCost: Retrieves the resource costs per second for territory
func (t *Territory) GetResourceCosts() map[ResourceType]int64 {
	costs := map[ResourceType]int64{
		CROP:    0,
		EMERALD: 0,
		FISH:    0,
		ORE:     0,
		WOOD:    0,
	}
	for k, v := range t.Bonuses {
		bonus := Bonuses[k]
		level := streams.StreamOf[BonusLevel](bonus.Levels...).Filter(func(e BonusLevel) bool {
			return e.Level == v
		}).Current[0]
		costs[bonus.UsedResorce] = costs[bonus.UsedResorce] + int64(level.Cost)
	}
	return costs
}

func (t *Territory) Tick() {
	currentTimeMillis := time.Now().UnixMilli()
	if currentTimeMillis-t.lastConsumedResource >= 1000 {
		t.lastConsumedResource = currentTimeMillis
		t.ConsumeResources(t.GetResourceCosts())
	}
	if currentTimeMillis-t.lastEmeraldProduced >= t.GetEmeraldRate()*1000 {
		t.lastEmeraldProduced = currentTimeMillis
		t.Storage[EMERALD] = t.Storage[EMERALD] + int64(t.GetProducedEmerald())
	}
	if currentTimeMillis-t.lastResourceProduced >= t.GetResourceRate()*1000 {
		t.lastEmeraldProduced = currentTimeMillis
		for k, v := range t.Storage {
			t.Storage[k] = v + int64(t.GetProducedResource())
		}
	}
	if currentTimeMillis-t.lastResourceTransfer >= 60000 {
		t.TransferResource(*t.Claim.GetTerritory(t.TargetTerritory).Wrapped)
	}
}

func (t *Territory) TransferResource(target Territory) {
	target.ReceiveResource(target, t.PassingResource)
}

func (t *Territory) ReceiveResource(target Territory, resources map[ResourceType]int64) {
	if t.Name == target.Name {
		t.StoreResource(resources)
	} else {
		t.PassingResource = resources
	}
}

func (t *Territory) ConsumeResources(costs map[ResourceType]int64) {
	for k, v := range costs {
		stored := t.Storage[k]
		if (stored - v) < 0 {
			passing := t.PassingResource[k]
			if passing+stored-v < 0 {
				t.resourceGap = true
			} else {
				t.Storage[k] = 0
				t.PassingResource[k] = passing + stored - v
				t.resourceGap = false
			}
		} else {
			t.resourceGap = false
			t.Storage[k] = t.Storage[k] - v
		}
	}
}

func (t *Territory) StoreResource(resources map[ResourceType]int64) {
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
