package config

import "fmt"

// ReadServiceConfig Reads Service params from config.json
func (v *viperConfig) ReadServiceConfig() string {
	return fmt.Sprintf("%s:%s", v.GetString("service.host"), v.GetString("service.port"))
}
