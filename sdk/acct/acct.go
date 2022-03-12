package acct

import (
	"strings"

	"random.chars.jp/git/misskey/sdk"
)

type Acct struct {
	Username string   `json:"username"`
	Host     sdk.Host `json:"host,omitempty"`
}

func (a *Acct) String() string {
	if a.Host.IsLocal() {
		return a.Username
	} else {
		return a.Username + "@" + string(a.Host)
	}
}

func Parse(str string) *Acct {
	if len(str) == 0 {
		return nil
	}

	if str[0] == '@' {
		str = str[1:]
	}

	s := strings.SplitN(str, "@", 2)
	switch len(s) {
	case 1:
		return &Acct{
			Username: s[0],
			Host:     "",
		}
	case 2:
		return &Acct{
			Username: s[0],
			Host:     (sdk.Host)(s[1]),
		}
	default:
		return nil
	}
}
