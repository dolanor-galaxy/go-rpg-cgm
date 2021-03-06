package combat

import (
	"github.com/steelx/go-rpg-cgm/world"
)

type ActorDef struct {
	Id               string //must match entityDef
	Stats            world.BaseStats
	StatGrowth       map[string]func() int
	Level            int
	ActionGrowth     map[int]map[string][]string //Level -> {Action : [special, special]}
	Portrait         string
	Name             string
	Actions          []string
	Magic            []string
	Special          []string
	StealItem        int //Item ID only for Enemy actors
	ActiveEquipSlots []int
	IsPlayer         bool
	Equipment
	Drop
}

type DropChanceItem struct {
	Oddment float64
	ItemId  int //item ID
}

type Drop struct {
	XP     float64
	Gold   [2]int //range min, max
	Always []int  //item ids that are guaranteed to drop
	Chance []DropChanceItem
}

type LevelUp struct {
	XP        float64
	Level     int
	BaseStats map[string]float64
	Actions   map[string][]string
}

//Must match to ItemsDB ID
type Equipment struct {
	Weapon,
	Armor,
	Access1,
	Access2 int //ItemsDB.Id
}
