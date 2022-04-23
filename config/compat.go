package config

import (
	"log"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
)

const confCompatDefaultPath = ".config/default.yml"

func (u *urlWrap) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	v, err := url.Parse(s)
	u.URL = v
	return err
}

type confCompat struct {
	URL  *urlWrap `yaml:"url"`
	Port int      `yaml:"port"`
	DB   struct {
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		DB           string `yaml:"db"`
		Username     string `yaml:"user"`
		Password     string `yaml:"pass"`
		DisableCache bool   `yaml:"disableCache,omitempty"`
		Extra        struct {
			SSL bool `yaml:"ssl,omitempty"`
		} `yaml:"extra,omitempty"`
	} `yaml:"db"`
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"pass,omitempty"`
		Prefix   string `yaml:"prefix,omitempty"`
		DB       int    `yaml:"db,omitempty"`
	} `yaml:"redis"`
	Elasticsearch *struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		SSL      bool   `yaml:"ssl"`
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"elasticsearch,omitempty"`
	IDGeneration          string `yaml:"id"`
	DisableHSTS           bool   `yaml:"disableHSTS,omitempty"`
	ClusterLimit          int    `yaml:"clusterLimit,omitempty"`
	DeliverJobConcurrency int    `yaml:"deliverJobConcurrency,omitempty"`
	InboxJobConcurrency   int    `yaml:"inboxJobConcurrency,omitempty"`
	DeliverJobPerSec      int    `yaml:"deliverJobPerSec,omitempty"`
	InboxJobPerSec        int    `yaml:"inboxJobPerSec,omitempty"`
	DeliverJobMaxAttempts int    `yaml:"deliverJobMaxAttempts,omitempty"`
	InboxJobMaxAttempts   int    `yaml:"inboxJobMaxAttempts,omitempty"`
	OutgoingAddressFamily string `yaml:"outgoingAddressFamily,omitempty"`
	Syslog                *struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"syslog,omitempty"`
	Proxy                  string   `yaml:"proxy,omitempty"`
	ProxyBypassHosts       []string `yaml:"proxyBypassHosts,omitempty"`
	ProxySmtp              string   `yaml:"proxySmtp,omitempty"`
	MediaProxy             string   `yaml:"mediaProxy,omitempty"`
	SignToActivityPubGet   bool     `yaml:"signToActivityPubGet,omitempty"`
	AllowedPrivateNetworks []string `yaml:"allowedPrivateNetworks,omitempty"`
	MaxFileSize            int      `yaml:"maxFileSize,omitempty"`
}

func tryCompat() bool {
	if loaded {
		panic("attempting compat parse after loaded set")
	}

	var c confCompat
	if file, err := os.Open(confCompatDefaultPath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatalf("error reading legacy yaml config file: %s", err)
	} else {
		if err = yaml.NewDecoder(file).Decode(&c); err != nil {
			log.Fatalf("error decoding legacy yaml config file: %s", err)
		}
	}
	Config = conf{
		System: systemConf{
			IDGeneration:         c.IDGeneration,
			MediaProxy:           c.MediaProxy,
			MaxFileSize:          c.MaxFileSize,
			SignToActivityPubGet: c.SignToActivityPubGet,
		},
		Web: webConf{
			URL:                 c.URL.URL,
			Host:                "127.0.0.1",
			Port:                c.Port,
			FastCGI:             false,
			ForwardedByClientIP: true,
			TrustedProxies:      c.AllowedPrivateNetworks,
			HSTS:                !c.DisableHSTS,
		},
		External: externalConf{
			Postgres: postgresConf{
				Host:     c.DB.Host,
				Port:     c.DB.Port,
				SSL:      c.DB.Extra.SSL,
				DB:       c.DB.DB,
				Username: c.DB.Username,
				Password: c.DB.Password,
				Cache:    !c.DB.DisableCache,
			},
			Redis: redisConf{
				Host:     c.Redis.Host,
				Port:     c.Redis.Port,
				DB:       c.Redis.DB,
				Prefix:   c.Redis.Prefix,
				Password: c.Redis.Password,
			},
		},
	}
	if c.Syslog != nil {
		Config.Log.Syslog = true
		Config.Log.Host = c.Syslog.Host
		Config.Log.Port = c.Syslog.Port
	}
	if c.Elasticsearch != nil {
		Config.External.Elasticsearch.Enable = true
		Config.External.Elasticsearch.Host = c.Elasticsearch.Host
		Config.External.Elasticsearch.Port = c.Elasticsearch.Port
		Config.External.Elasticsearch.SSL = c.Elasticsearch.SSL
		Config.External.Elasticsearch.Username = c.Elasticsearch.Username
		Config.External.Elasticsearch.Password = c.Elasticsearch.Password
	}
	if c.MaxFileSize == 0 {
		Config.System.MaxFileSize = defConf.System.MaxFileSize
	}
	return true
}
