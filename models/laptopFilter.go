package models

type LaptopFilter struct {
	ID                uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	LaptopID          uint    `json:"laptopId"`
	ProcessorBrand    string  `gorm:"not null" json:"processorBrand"`
	ProcessorName     string  `gorm:"not null" json:"processorName"`
	GraphicsCardType  string  `gorm:"not null" json:"graphicsCardType"`
	GraphicsCardName  string  `gorm:"not null" json:"graphicsCardName"`
	RAM               int     `gorm:"not null" json:"ram"`
	Storage           int     `gorm:"not null" json:"storage"`
	BatteryCapacity   int     `gorm:"not null" json:"batteryCapacity"`
	OperatingSystem   string  `gorm:"not null" json:"operatingSystem"`
	ScreenSize        float64 `gorm:"not null" json:"screenSize"`
	ScreenRefreshRate int     `gorm:"not null" json:"screenRefreshRate"`
	ScreenType        string  `gorm:"not null" json:"screenType"`
	BodyMaterial      string  `gorm:"not null" json:"bodyMaterial"`
	Width             float64 `gorm:"not null" json:"width"`
	Height            float64 `gorm:"not null" json:"height"`
	Depth             float64 `gorm:"not null" json:"depth"`
	Weight            float64 `gorm:"not null" json:"weight"`
}
