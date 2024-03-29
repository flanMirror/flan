package config

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var (
	Log      logConf
	System   systemConf
	Web      webConf
	Proxy    *proxyConf
	External externalConf
)

var (
	HTTPS bool
	URL   string
)

const ConfDefaultPath = "misskey.conf"

type urlWrap struct {
	*url.URL
}

type conf struct {
	Log      logConf      `toml:"log"`
	System   systemConf   `toml:"system"`
	Web      webConf      `toml:"web"`
	Proxy    *proxyConf   `toml:"proxy,omitempty"`
	External externalConf `toml:"external"`
}

type logConf struct {
	Verbose bool   `toml:"verbose"`
	Syslog  bool   `toml:"syslog"`
	Host    string `toml:"host"`
	Port    int    `toml:"port"`
}

type systemConf struct {
	IDGeneration string `toml:"id_generation"`
	// TODO: evaluate whether this is necessary considering go has built-in concurrency
	//ClusterLimit          int    `yaml:"clusterLimit"`
	//DeliverJobConcurrency int    `yaml:"deliverJobConcurrency"`
	//InboxJobConcurrency   int    `yaml:"inboxJobConcurrency"`
	//DeliverJobPerSec      int    `yaml:"deliverJobPerSec"`
	//InboxJobPerSec        int    `yaml:"inboxJobPerSec"`
	//DeliverJobMaxAttempts int    `yaml:"deliverJobMaxAttempts"`
	//InboxJobMaxAttempts   int    `yaml:"inboxJobMaxAttempts"`
	MediaProxy           string `toml:"media_proxy"`
	MaxFileSize          int    `toml:"max_file_size"`
	SignToActivityPubGet bool   `toml:"sign_to_activity_pub_get"`
}

type webConf struct {
	URL                 *url.URL `toml:"url"`
	Host                string   `toml:"host"`
	Port                int      `toml:"port"`
	FastCGI             bool     `toml:"fcgi"`
	ForwardedByClientIP bool     `toml:"proxy"`
	TrustedProxies      []string `json:"trusted_proxies"`
	HSTS                bool     `toml:"hsts"`
}

type proxyConf struct {
	Address     string   `toml:"address"`
	BypassHosts []string `toml:"bypass_hosts"`
	SMTP        string   `toml:"smtp"`
}

type externalConf struct {
	Postgres      postgresConf      `toml:"postgres"`
	Redis         redisConf         `toml:"redis"`
	Elasticsearch elasticsearchConf `toml:"elasticsearch"`
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
	Enable   bool   `toml:"enable"`
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
	verbose  bool
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, "enable verbose logging")
}

func Parse() bool {
	defer func() {
		loaded = true
		if verbose {
			Config.Log.Verbose = true
		}
		Log = Config.Log
		System = Config.System
		Web = Config.Web
		Proxy = Config.Proxy
		External = Config.External

		/* misskey upstream checks whether the https section in config file is set for this
		which is absolutely insane since no sane person would run anything like that as root,
		this is clearly a bug, so we're doing it the proper way here */
		HTTPS = Web.URL.Scheme == "https"

		// necessary to avoid having two /
		URL = strings.TrimSuffix(Config.Web.URL.String(), "/")
	}()
	if loaded {
		panic("attempting config parse after loaded set")
	}

	if meta, err := toml.DecodeFile(ConfPath, &Config); err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("error parsing configuration: %s", err)
		}

		if tryCompat() {
			return false
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

	if Config.Web.URL == nil {
		log.Fatalf("no URL specified")
	}

	// set up proxies here for the default http client
	if Proxy.Address != "" {
		if u, err := url.Parse(Proxy.Address); err != nil {
			log.Fatalf("proxy address present but invalid: %s", err)
		} else {
			bypass := make(map[string]bool, len(Proxy.BypassHosts))
			for _, host := range Proxy.BypassHosts {
				bypass[host] = true
			}
			http.DefaultClient.Transport = &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
				if bypass[req.Host] {
					return nil, nil
				}
				return u, nil
			}}
		}
	}

	return true
}

var defConf = conf{
	Log: logConf{
		Verbose: false,
		Syslog:  false,
		Host:    "localhost",
		Port:    514,
	},
	System: systemConf{
		IDGeneration:         "aid",
		MediaProxy:           "",
		MaxFileSize:          268435456,
		SignToActivityPubGet: true,
	},
	Web: webConf{
		URL:                 nil,
		Host:                "127.0.0.1",
		Port:                3000,
		FastCGI:             true,
		ForwardedByClientIP: true,
		TrustedProxies:      []string{"127.0.0.1/32", "172.16.0.0/12", "10.0.0.0/8"},
		HSTS:                false,
	},
	External: externalConf{
		Postgres: postgresConf{
			Host:     "/tmp",
			Port:     5432,
			SSL:      false,
			DB:       "misskey",
			Username: "misskey",
			Password: "",
			Cache:    false,
		},
		Redis: redisConf{
			Host:     "localhost",
			Port:     6379,
			DB:       9,
			Prefix:   "",
			Password: "",
		},
		Elasticsearch: elasticsearchConf{
			Enable:   false,
			Host:     "localhost",
			Port:     9200,
			SSL:      false,
			Username: "",
			Password: "",
		},
	},
}
