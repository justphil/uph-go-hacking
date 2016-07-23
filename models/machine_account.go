package models

type MachineAccount struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	OsUser     string `json:"osUser"`
	OsPassword string `json:"osPassword"`
}
