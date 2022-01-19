package config

const confCompatDefaultPath = ".config/default.yml"

type confCompat struct {
	URL  string `yaml:"url"`
	Port int    `yaml:"port"`
	DB   struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		DB           string `yaml:"db"`
		Username     string `yaml:"user"`
		Password     string `yaml:"pass"`
		DisableCache bool   `yaml:"disableCache"`
		Extra        struct {
			SSL bool `yaml:"ssl"`
		} `yaml:"extra"`
	} `yaml:"db"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"pass"`
		Prefix   string `yaml:"prefix"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	Elasticsearch struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		SSL      bool   `yaml:"ssl"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"elasticsearch"`
	IDGeneration          string `yaml:"id"`
	DisableHSTS           bool   `yaml:"disableHSTS"`
	ClusterLimit          int    `yaml:"clusterLimit"`
	DeliverJobConcurrency int    `yaml:"deliverJobConcurrency"`
	InboxJobConcurrency   int    `yaml:"inboxJobConcurrency"`
	DeliverJobPerSec      int    `yaml:"deliverJobPerSec"`
	InboxJobPerSec        int    `yaml:"inboxJobPerSec"`
	DeliverJobMaxAttempts int    `yaml:"deliverJobMaxAttempts"`
	InboxJobMaxAttempts   int    `yaml:"inboxJobMaxAttempts"`
	OutgoingAddressFamily string `yaml:"outgoingAddressFamily"`
	Syslog                struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"syslog"`
	Proxy                  string   `yaml:"proxy"`
	ProxyBypassHosts       []string `yaml:"proxyBypassHosts"`
	ProxySmtp              string   `yaml:"proxySmtp"`
	MediaProxy             string   `yaml:"mediaProxy"`
	SignToActivityPubGet   bool     `yaml:"signToActivityPubGet"`
	AllowedPrivateNetworks []string `yaml:"allowedPrivateNetworks"`
	MaxFileSize            int      `yaml:"maxFileSize"`
}

func tryCompat() bool {
	if loaded {
		panic("attempting compat parse after loaded set")
	}

	// TODO: parse
	// TODO: defaults
	return false
}
