package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
func adminAbuseUserReports(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAccountsCreate - admin/accounts/create
func adminAccountsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAccountsDelete - admin/accounts/delete
func adminAccountsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAdCreate - admin/ad/create
func adminAdCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAdDelete - admin/ad/delete
func adminAdDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAdList - admin/ad/list
func adminAdList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAdUpdate - admin/ad/update
func adminAdUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAnnouncementsCreate - admin/announcements/create
func adminAnnouncementsCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAnnouncementsDelete - admin/announcements/delete
func adminAnnouncementsDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAnnouncementsList - admin/announcements/list
func adminAnnouncementsList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminAnnouncementsUpdate - admin/announcements/update
func adminAnnouncementsUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminDeleteAllFilesOfAUser - admin/delete-all-files-of-a-user
func adminDeleteAllFilesOfAUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminDeleteLogs - admin/delete-logs
func adminDeleteLogs(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminDriveCleanRemoteFiles - admin/drive/clean-remote-files
func adminDriveCleanRemoteFiles(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminDriveCleanup - admin/drive/cleanup
func adminDriveCleanup(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminDriveFiles - admin/drive/files
func adminDriveFiles(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminDriveShowFile - admin/drive/show-file
func adminDriveShowFile(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminEmojiAdd - admin/emoji/add
func adminEmojiAdd(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminEmojiCopy - admin/emoji/copy
func adminEmojiCopy(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminEmojiList - admin/emoji/list
func adminEmojiList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminEmojiListRemote - admin/emoji/list-remote
func adminEmojiListRemote(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminEmojiRemove - admin/emoji/remove
func adminEmojiRemove(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminEmojiUpdate - admin/emoji/update
func adminEmojiUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminFederationDeleteAllFiles - admin/federation/delete-all-files
func adminFederationDeleteAllFiles(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminFederationRefreshRemoteInstanceMetadata - admin/federation/refresh-remote-instance-metadata
func adminFederationRefreshRemoteInstanceMetadata(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminFederationRemoveAllFollowing - admin/federation/remove-all-following
func adminFederationRemoveAllFollowing(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminFederationUpdateInstance - admin/federation/update-instance
func adminFederationUpdateInstance(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminGetIndexStats - admin/get-index-stats
func adminGetIndexStats(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminGetTableStats - admin/get-table-stats
func adminGetTableStats(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminInvite - admin/invite
func adminInvite(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminModeratorsAdd - admin/moderators/add
func adminModeratorsAdd(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminModeratorsRemove - admin/moderators/remove
func adminModeratorsRemove(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminPromoCreate - admin/promo/create
func adminPromoCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminQueueClear - admin/queue/clear
func adminQueueClear(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminQueueDeliverDelayed - admin/queue/deliver-delayed
func adminQueueDeliverDelayed(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminQueueInboxDelayed - admin/queue/inbox-delayed
func adminQueueInboxDelayed(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminQueueJobs - admin/queue/jobs
func adminQueueJobs(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminQueueStats - admin/queue/stats
func adminQueueStats(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminRelaysAdd - admin/relays/add
func adminRelaysAdd(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminRelaysList - admin/relays/list
func adminRelaysList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminRelaysRemove - admin/relays/remove
func adminRelaysRemove(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminResetPassword - admin/reset-password
func adminResetPassword(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminResolveAbuseUserReport - admin/resolve-abuse-user-report
func adminResolveAbuseUserReport(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminResyncChart - admin/resync-chart
func adminResyncChart(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminSendEmail - admin/send-email
func adminSendEmail(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminServerInfo - admin/server-info
func adminServerInfo(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminShowModerationLogs - admin/show-moderation-logs
func adminShowModerationLogs(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminShowUser - admin/show-user
func adminShowUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminShowUsers - admin/show-users
func adminShowUsers(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminSilenceUser - admin/silence-user
func adminSilenceUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminSuspendUser - admin/suspend-user
func adminSuspendUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminUnsilenceUser - admin/unsilence-user
func adminUnsilenceUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminUnsuspendUser - admin/unsuspend-user
func adminUnsuspendUser(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminUpdateMeta - admin/update-meta
func adminUpdateMeta(c *gin.Context) {
	// TODO
	placeholder(c)
}

// adminVacuum - admin/vacuum
func adminVacuum(c *gin.Context) {
	// TODO
	placeholder(c)
}
