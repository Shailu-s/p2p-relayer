package config

type NodeDetailsConfig struct {
	WorkerAddress string
	PrivateKey    string
}

type P2PConfig struct {
	Port           int64
	ListenAddress  string
	DiscoveryTag   string
	RelayerAddress []string
}

func (v *viperConfig) ReadNodeDetailsConfig() NodeDetailsConfig {
	return NodeDetailsConfig{
		WorkerAddress: v.GetString("node_details.worker_addr"),
		PrivateKey:    v.GetString("node_details.private_key"),
	}
}

func (v *viperConfig) ReadP2PConfig() P2PConfig {
	return P2PConfig{
		Port:           v.GetInt64("p2p_config.port"),
		ListenAddress:  v.GetString("p2p_config.listen_address"),
		DiscoveryTag:   v.GetString("p2p_config.discovery_tag"),
		RelayerAddress: v.GetStringSlice("p2p_config.relayer_address"),
	}
}
