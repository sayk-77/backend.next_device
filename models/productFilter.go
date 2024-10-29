package models

type ProductFilter struct {
	ID            uint    `gorm:"primaryKey"`
	ProductID     uint    `gorm:"not null"`
	DisplaySize   float64 `gorm:"type:decimal(10,2)"`
	RAM           uint    `gorm:"not null"`
	Storage       uint    `gorm:"not null"`
	CameraQuality uint    `gorm:"not null"`
	Processor     string  `gorm:"size:100"`
	Battery       uint    `gorm:"not null"`
	OS            string  `gorm:"size:50"`
	Width         float64 `gorm:"type:decimal(10,2)"`
	Height        float64 `gorm:"type:decimal(10,2)"`
	Length        float64 `gorm:"type:decimal(10,2)"`
	Weight        float64 `gorm:"type:decimal(10,2)"`
}
