package pkg

import "time"

const (
	CONFIG_FILE               string        = "config.json"
	PORT                      int           = 9090
	CONNECTION_STRING_DEFAULT string        = "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	TRANS_TIMEOUT             time.Duration = 5 * time.Second
)
