package qhelper

import (
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var NotSuspended = qm.Where(`"isSuspended" = false`)

func UsernameLower(username string) qm.QueryMod {
	return qm.Where(`"usernameLower" = ?`, strings.ToLower(username))
}
