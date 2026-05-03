package config

type Config struct {
	Address     string `json:"address"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DbName      string `json:"database"`
	ServiceType string `json:"serviceType"`
}
