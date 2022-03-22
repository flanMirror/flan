package qhelper

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

var (
	NotRenote         = qm.Where(`"renoteId" is null`)
	Renote            = qm.Where(`"renoteId" is not null`)
	PubliclyAvailable = qm.Where(`visibility in ('public', 'home')`)
)

func UserID(id string) qm.QueryMod {
	return qm.Where(`"userId" = ?`, id)
}
