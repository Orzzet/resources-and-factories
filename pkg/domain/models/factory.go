package models

type FactoryType string

const (
	IronFactory   FactoryType = "ironFactory"
	CopperFactory FactoryType = "copperFactory"
	GoldFactory   FactoryType = "goldFactory"
)

type FactoriesStats map[FactoryType]map[int]FactoryStats

type FactoryStats struct {
	FactoryType
	Level               uint
	Production          Resource
	NextUpgradeDuration uint
	NextUpgradeCost     []Resource
}

type Factory struct {
	FactoryStats
	FactoryType
	Level            uint
	IsUpdating       bool
	UpgradeStartedIn uint
}

func GetFactoriesStats() FactoriesStats {
	factoriesStats := make(FactoriesStats)
	factoriesStats[IronFactory] = map[int]FactoryStats{
		1: {
			Level:               1,
			Production:          Resource{IronResource, 10},
			NextUpgradeDuration: 15,
			NextUpgradeCost: []Resource{
				{IronResource, 300}, {CopperResource, 100}, {GoldResource, 1},
			},
		},
		2: {
			Level:               2,
			Production:          Resource{IronResource, 20},
			NextUpgradeDuration: 30,
			NextUpgradeCost: []Resource{
				{IronResource, 800}, {CopperResource, 250}, {GoldResource, 2},
			},
		},
	}
	factoriesStats[CopperFactory] = map[int]FactoryStats{
		1: {
			Level:               1,
			Production:          Resource{CopperResource, 3},
			NextUpgradeDuration: 15,
			NextUpgradeCost: []Resource{
				{IronResource, 200}, {CopperResource, 70},
			},
		},
		2: {
			Level:               2,
			Production:          Resource{CopperResource, 7},
			NextUpgradeDuration: 30,
			NextUpgradeCost: []Resource{
				{IronResource, 400}, {CopperResource, 150},
			},
		},
	}
	factoriesStats[GoldFactory] = map[int]FactoryStats{
		1: {
			Level:               1,
			Production:          Resource{GoldResource, 2 / 60},
			NextUpgradeDuration: 15,
			NextUpgradeCost: []Resource{
				{CopperResource, 100}, {GoldResource, 2},
			},
		},
		2: {
			Level:               2,
			Production:          Resource{GoldResource, 2 / 60},
			NextUpgradeDuration: 30,
			NextUpgradeCost: []Resource{
				{CopperResource, 200}, {GoldResource, 4},
			},
		},
	}
	return factoriesStats
}
