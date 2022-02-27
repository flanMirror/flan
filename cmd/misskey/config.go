package main

import (
	"flag"

	"random.chars.jp/git/misskey/config"
)

func init() {
	flag.StringVar(&config.ConfPath, "c", config.ConfDefaultPath, "path to configuration file")
}
