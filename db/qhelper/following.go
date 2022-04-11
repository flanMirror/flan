package qhelper

import "github.com/volatiletech/sqlboiler/v4/queries/qm"

func FolloweeID(id string) qm.QueryMod {
	return qm.Where(`"followeeId" = ?`, id)
}

func FollowerID(id string) qm.QueryMod {
	return qm.Where(`"followerId" = ?`, id)
}
