package qhelper

import (
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"random.chars.jp/git/misskey/sdk"
)

var LocalHost = qm.Where("host is null")

func Host(host sdk.Host) qm.QueryMod {
	if host.IsLocal() {
		return LocalHost
	} else {
		return qm.Where("host = ?", host)
	}
}
