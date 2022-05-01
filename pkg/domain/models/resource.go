package models

type ResourceType string

const (
	IronResource   ResourceType = "ironResource"
	CopperResource ResourceType = "copperResource"
	GoldResource   ResourceType = "goldResource"
)

type Resource struct {
	ResourceType
	Amount float64
}
