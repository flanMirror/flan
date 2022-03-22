package qhelper

import (
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"random.chars.jp/git/misskey/db"
)

func ExpiresAt(t time.Time) qm.QueryMod {
	return qm.Where(
		`"expiresAt" >= (? at time zone 'utc')`,
		t.UTC().Format(db.TimestampFormat),
	)
}
