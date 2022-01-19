package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

var (
	Log      logConf
	System   systemConf
	Web      webConf
	External externalConf
)

const ConfDefaultPath = "misskey.conf"

type conf struct {
	Log      logConf      `toml:"log"`
	System   systemConf   `toml:"system"`
	Web      webConf      `toml:"web"`
	External externalConf `toml:"external"`
}

type logConf struct {
	Verbose bool   `toml:"verbose"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
}

type systemConf struct {
	IDGeneration string `toml:"id_generation"`
	// TODO: map these
	//ClusterLimit          int    `yaml:"clusterLimit"`
	//DeliverJobConcurrency int    `yaml:"deliverJobConcurrency"`
	//InboxJobConcurrency   int    `yaml:"inboxJobConcurrency"`
	//DeliverJobPerSec      int    `yaml:"deliverJobPerSec"`
	//InboxJobPerSec        int    `yaml:"inboxJobPerSec"`
	//DeliverJobMaxAttempts int    `yaml:"deliverJobMaxAttempts"`
	//InboxJobMaxAttempts   int    `yaml:"inboxJobMaxAttempts"`
	//Proxy                  string   `yaml:"proxy"`
	//ProxyBypassHosts       []string `yaml:"proxyBypassHosts"`
	//ProxySmtp              string   `yaml:"proxySmtp"`
	MediaProxy           string `toml:"media_proxy"`
	MaxFileSize          int    `toml:"max_file_size"`
	SignToActivityPubGet bool   `toml:"sign_to_activity_pub_get"`
}

type webConf struct {
	URL                 string   `toml:"url"`
	Host                string   `toml:"host"`
	Port                int      `toml:"port"`
	ForwardedByClientIP bool     `toml:"proxy"`
	TrustedProxies      []string `json:"trusted_proxies"`
	HSTS                bool     `toml:"hsts"`
}

type externalConf struct {
	Postgres      postgresConf      `toml:"postgres"`
	Redis         redisConf         `toml:"redis"`
	elasticsearch elasticsearchConf `toml:"elasticsearch"`
}

type postgresConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	SSL      bool   `toml:"ssl"`
	DB       string `toml:"db"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Cache    bool   `toml:"cache"`
}

type redisConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DB       int    `toml:"db"`
	Prefix   string `toml:"prefix"`
	Password string `toml:"password"`
}

type elasticsearchConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	SSL      bool   `toml:"ssl"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

var (
	ConfPath string
	Config   = conf{}
	loaded   = false
)

func Parse() {
	defer func() {
		loaded = true
		Log = Config.Log
		System = Config.System
		Web = Config.Web
		External = Config.External
	}()
	if loaded {
		panic("attempting config parse after loaded set")
	}

	if meta, err := toml.DecodeFile(ConfPath, &Config); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("error parsing configuration: %s", err)
		}

		if tryCompat() {
			return
		}

		var file *os.File
		if file, err = os.Create(ConfPath); err != nil {
			log.Fatalf("error creating configuration file: %s", err)
		} else if err = toml.NewEncoder(file).Encode(defConf); err != nil {
			log.Fatalf("error generating default configuration: %s", err)
		}
		log.Fatal("configuration file with default values generated")
	} else {
		for _, key := range meta.Undecoded() {
			log.Printf("unused key in configuration file: %s", key.String())
		}
	}
}

var defConf = conf{
	// TODO
}
