package manager

import (
	"encoding/json"
	"os"

	"dev.justdrven/loadbalancer/data"
)

const CONFIG_FILE = "config.json"

func CCreateConfig() (error, *data.Config) {
	bytes, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		return err, nil
	}

	var context data.Config
	jsonErr := json.Unmarshal(bytes, &context)

	return jsonErr, &context
}
