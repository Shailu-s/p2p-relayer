package config

// MatchingConfig stores required config for matching critirea and p2p
type MatchingConfig struct {
	Leader         bool  `json:"leader"`
	MaxFailAllowed int64 `json:"max_fail_allowed"`
}

// ReadMatchingConfig reads matching params from config.json
func (v *viperConfig) ReadMatchingConfig() MatchingConfig {
	return MatchingConfig{
		Leader:         v.GetBool("matching_config.leader"),
		MaxFailAllowed: v.GetInt64("matching_config.max_fail_allowed"),
	}
}

// ReadLogLevelConfig reads logger level from config.json
func (v *viperConfig) ReadLogLevelConfig() string {
	return v.GetString("logger-level")
}
