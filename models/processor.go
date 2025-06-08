package models

type Processor struct {
	ID        uint    `json: "id" gorm:"primarykey"`
	Name      string `json:"name"`
	Cores     int    `json:"cores"`
	Frequency string `json:"frequency"`
}
