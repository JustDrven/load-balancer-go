package config

import (
	"encoding/json"
	"os"

	"dev.justdrven/loadbalancer/pkg"
)

func CreateConfig() (error, *Config) {
	bytes, err := os.ReadFile(pkg.CONFIG_FILE)
	if err != nil {
		return err, nil
	}

	var context Config
	jsonErr := json.Unmarshal(bytes, &context)

	return jsonErr, &context
}
