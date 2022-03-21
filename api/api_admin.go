package api

import (
	"net/http"

	"random.chars.jp/git/misskey/data/payload"
	"random.chars.jp/git/misskey/db/orm"
)

func init() {
	route{http.MethodPost, "admin/abuse-user-reports", adminAbuseUserReports}.register()
	route{http.MethodPost, "admin/accounts/create", adminAccountsCreate}.register()
	route{http.MethodPost, "admin/accounts/delete", adminAccountsDelete}.register()
	route{http.MethodPost, "admin/ad/create", adminAdCreate}.register()
	route{http.MethodPost, "admin/ad/delete", adminAdDelete}.register()
	route{http.MethodPost, "admin/ad/list", adminAdList}.register()
	route{http.MethodPost, "admin/ad/update", adminAdUpdate}.register()
	route{http.MethodPost, "admin/announcements/create", adminAnnouncementsCreate}.register()
	route{http.MethodPost, "admin/announcements/delete", adminAnnouncementsDelete}.register()
	route{http.MethodPost, "admin/announcements/list", adminAnnouncementsList}.register()
	route{http.MethodPost, "admin/announcements/update", adminAnnouncementsUpdate}.register()
	route{http.MethodPost, "admin/delete-all-files-of-a-user", adminDeleteAllFilesOfAUser}.register()
	route{http.MethodPost, "admin/delete-logs", adminDeleteLogs}.register()
	route{http.MethodPost, "admin/drive/clean-remote-files", adminDriveCleanRemoteFiles}.register()
	route{http.MethodPost, "admin/drive/cleanup", adminDriveCleanup}.register()
	route{http.MethodPost, "admin/drive/files", adminDriveFiles}.register()
	route{http.MethodPost, "admin/drive/show-file", adminDriveShowFile}.register()
	route{http.MethodPost, "admin/emoji/add", adminEmojiAdd}.register()
	route{http.MethodPost, "admin/emoji/copy", adminEmojiCopy}.register()
	route{http.MethodPost, "admin/emoji/list", adminEmojiList}.register()
	route{http.MethodPost, "admin/emoji/list-remote", adminEmojiListRemote}.register()
	route{http.MethodPost, "admin/emoji/remove", adminEmojiRemove}.register()
	route{http.MethodPost, "admin/emoji/update", adminEmojiUpdate}.register()
	route{http.MethodPost, "admin/federation/delete-all-files", adminFederationDeleteAllFiles}.register()
	route{http.MethodPost, "admin/federation/refresh-remote-instance-metadata", adminFederationRefreshRemoteInstanceMetadata}.register()
	route{http.MethodPost, "admin/federation/remove-all-following", adminFederationRemoveAllFollowing}.register()
	route{http.MethodPost, "admin/federation/update-instance", adminFederationUpdateInstance}.register()
	route{http.MethodPost, "admin/get-index-stats", adminGetIndexStats}.register()
	route{http.MethodPost, "admin/get-table-stats", adminGetTableStats}.register()
	route{http.MethodPost, "admin/invite", adminInvite}.register()
	route{http.MethodPost, "admin/moderators/add", adminModeratorsAdd}.register()
	route{http.MethodPost, "admin/moderators/remove", adminModeratorsRemove}.register()
	route{http.MethodPost, "admin/promo/create", adminPromoCreate}.register()
	route{http.MethodPost, "admin/queue/clear", adminQueueClear}.register()
	route{http.MethodPost, "admin/queue/deliver-delayed", adminQueueDeliverDelayed}.register()
	route{http.MethodPost, "admin/queue/inbox-delayed", adminQueueInboxDelayed}.register()
	route{http.MethodPost, "admin/queue/jobs", adminQueueJobs}.register()
	route{http.MethodPost, "admin/queue/stats", adminQueueStats}.register()
	route{http.MethodPost, "admin/relays/add", adminRelaysAdd}.register()
	route{http.MethodPost, "admin/relays/list", adminRelaysList}.register()
	route{http.MethodPost, "admin/relays/remove", adminRelaysRemove}.register()
	route{http.MethodPost, "admin/reset-password", adminResetPassword}.register()
	route{http.MethodPost, "admin/resolve-abuse-user-report", adminResolveAbuseUserReport}.register()
	route{http.MethodPost, "admin/resync-chart", adminResyncChart}.register()
	route{http.MethodPost, "admin/send-email", adminSendEmail}.register()
	route{http.MethodPost, "admin/server-info", adminServerInfo}.register()
	route{http.MethodPost, "admin/show-moderation-logs", adminShowModerationLogs}.register()
	route{http.MethodPost, "admin/show-user", adminShowUser}.register()
	route{http.MethodPost, "admin/show-users", adminShowUsers}.register()
	route{http.MethodPost, "admin/silence-user", adminSilenceUser}.register()
	route{http.MethodPost, "admin/suspend-user", adminSuspendUser}.register()
	route{http.MethodPost, "admin/unsilence-user", adminUnsilenceUser}.register()
	route{http.MethodPost, "admin/unsuspend-user", adminUnsuspendUser}.register()
	route{http.MethodPost, "admin/update-meta", adminUpdateMeta}.register()
	route{http.MethodPost, "admin/vacuum", adminVacuum}.register()
}

// adminAbuseUserReports - admin/abuse-user-reports
func adminAbuseUserReports(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAccountsCreate - admin/accounts/create
func adminAccountsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAccountsDelete - admin/accounts/delete
func adminAccountsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAdCreate - admin/ad/create
func adminAdCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAdDelete - admin/ad/delete
func adminAdDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAdList - admin/ad/list
func adminAdList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAdUpdate - admin/ad/update
func adminAdUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAnnouncementsCreate - admin/announcements/create
func adminAnnouncementsCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAnnouncementsDelete - admin/announcements/delete
func adminAnnouncementsDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAnnouncementsList - admin/announcements/list
func adminAnnouncementsList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminAnnouncementsUpdate - admin/announcements/update
func adminAnnouncementsUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminDeleteAllFilesOfAUser - admin/delete-all-files-of-a-user
func adminDeleteAllFilesOfAUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminDeleteLogs - admin/delete-logs
func adminDeleteLogs(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminDriveCleanRemoteFiles - admin/drive/clean-remote-files
func adminDriveCleanRemoteFiles(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminDriveCleanup - admin/drive/cleanup
func adminDriveCleanup(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminDriveFiles - admin/drive/files
func adminDriveFiles(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminDriveShowFile - admin/drive/show-file
func adminDriveShowFile(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminEmojiAdd - admin/emoji/add
func adminEmojiAdd(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminEmojiCopy - admin/emoji/copy
func adminEmojiCopy(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminEmojiList - admin/emoji/list
func adminEmojiList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminEmojiListRemote - admin/emoji/list-remote
func adminEmojiListRemote(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminEmojiRemove - admin/emoji/remove
func adminEmojiRemove(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminEmojiUpdate - admin/emoji/update
func adminEmojiUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminFederationDeleteAllFiles - admin/federation/delete-all-files
func adminFederationDeleteAllFiles(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminFederationRefreshRemoteInstanceMetadata - admin/federation/refresh-remote-instance-metadata
func adminFederationRefreshRemoteInstanceMetadata(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminFederationRemoveAllFollowing - admin/federation/remove-all-following
func adminFederationRemoveAllFollowing(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminFederationUpdateInstance - admin/federation/update-instance
func adminFederationUpdateInstance(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminGetIndexStats - admin/get-index-stats
func adminGetIndexStats(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminGetTableStats - admin/get-table-stats
func adminGetTableStats(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminInvite - admin/invite
func adminInvite(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminModeratorsAdd - admin/moderators/add
func adminModeratorsAdd(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminModeratorsRemove - admin/moderators/remove
func adminModeratorsRemove(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminPromoCreate - admin/promo/create
func adminPromoCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminQueueClear - admin/queue/clear
func adminQueueClear(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminQueueDeliverDelayed - admin/queue/deliver-delayed
func adminQueueDeliverDelayed(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminQueueInboxDelayed - admin/queue/inbox-delayed
func adminQueueInboxDelayed(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminQueueJobs - admin/queue/jobs
func adminQueueJobs(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminQueueStats - admin/queue/stats
func adminQueueStats(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminRelaysAdd - admin/relays/add
func adminRelaysAdd(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminRelaysList - admin/relays/list
func adminRelaysList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminRelaysRemove - admin/relays/remove
func adminRelaysRemove(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminResetPassword - admin/reset-password
func adminResetPassword(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminResolveAbuseUserReport - admin/resolve-abuse-user-report
func adminResolveAbuseUserReport(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminResyncChart - admin/resync-chart
func adminResyncChart(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminSendEmail - admin/send-email
func adminSendEmail(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminServerInfo - admin/server-info
func adminServerInfo(ctx Context) {
	ctx.RequireCredential(func(ctx Context, user *orm.User) {
		if !user.IsModerator && !user.IsAdmin {
			ctx.RawJSON(http.StatusForbidden, payload.AccessDeniedNotModerator.Data())
			return
		} else {
			ctx.RawJSON(http.StatusOK, payload.AdminServerInfo.Data())
		}
	})
}

// adminShowModerationLogs - admin/show-moderation-logs
func adminShowModerationLogs(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminShowUser - admin/show-user
func adminShowUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminShowUsers - admin/show-users
func adminShowUsers(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminSilenceUser - admin/silence-user
func adminSilenceUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminSuspendUser - admin/suspend-user
func adminSuspendUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminUnsilenceUser - admin/unsilence-user
func adminUnsilenceUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminUnsuspendUser - admin/unsuspend-user
func adminUnsuspendUser(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminUpdateMeta - admin/update-meta
func adminUpdateMeta(ctx Context) {
	// TODO
	placeholder(ctx)
}

// adminVacuum - admin/vacuum
func adminVacuum(ctx Context) {
	// TODO
	placeholder(ctx)
}
