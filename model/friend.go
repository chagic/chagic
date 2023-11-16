package model

type Friend struct {
	Base
	SourceID string `gorm:"column:source_id"`
	Source   User   `gorm:"foreignKey:SourceID"`
	TargetID string `gorm:"column:target_id"`
	Target   User   `gorm:"foreignKey:TargetID"`
}
