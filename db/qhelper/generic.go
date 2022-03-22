package qhelper

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

func IDString(id string) qm.QueryMod {
	return qm.Where(`"id" = ?`, id)
}

func ID(id string) qm.QueryMod {
	return qm.Where(`id = ?`, id)
}
