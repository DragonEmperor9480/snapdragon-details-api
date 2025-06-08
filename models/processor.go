package models

type Processor struct {
	ID        int    `json: "id"`
	Name      string `json:"name"`
	Cores     int    `json:"cores"`
	Frequency string `json:"frequency"`
}
