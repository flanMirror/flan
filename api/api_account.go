package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{
		Method:  http.MethodPost,
		Pattern: "blocking/create",
		Handler: blockingCreate,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "blocking/delete",
		Handler: blockingDelete,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "blocking/list",
		Handler: blockingList,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "clips/add-note",
		Handler: clipsAddNote,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "clips/notes",
		Handler: clipsNotes,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i",
		Handler: i,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/favorites",
		Handler: iFavorites,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/gallery/likes",
		Handler: iGalleryLikes,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/gallery/posts",
		Handler: iGalleryPosts,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/get-word-muted-notes-count",
		Handler: iGetWordMutedNotesCount,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/notifications",
		Handler: iNotifications,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/page-likes",
		Handler: iPageLikes,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/pages",
		Handler: iPages,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/pin",
		Handler: iPin,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/read-all-messaging-messages",
		Handler: iReadAllMessagingMessages,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/read-all-unread-notes",
		Handler: iReadAllUnreadNotes,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/read-announcement",
		Handler: iReadAnnouncement,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/unpin",
		Handler: iUnpin,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/update",
		Handler: iUpdate,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "i/user-group-invites",
		Handler: iUserGroupInvites,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "mute/create",
		Handler: muteCreate,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "mute/delete",
		Handler: muteDelete,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "mute/list",
		Handler: muteList,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "my/apps",
		Handler: myApps,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "sw/register",
		Handler: swRegister,
	}.register()

	route{
		Method:  http.MethodPost,
		Pattern: "sw/unregister",
		Handler: swUnregister,
	}.register()
}

// blockingCreate - blocking/create
func blockingCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// blockingDelete - blocking/delete
func blockingDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// blockingList - blocking/list
func blockingList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// clipsAddNote - clips/add-note
func clipsAddNote(c *gin.Context) {
	// TODO
	placeholder(c)
}

// clipsNotes - clips/notes
func clipsNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// i - i
func i(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iFavorites - i/favorites
func iFavorites(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iGalleryLikes - i/gallery/likes
func iGalleryLikes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iGalleryPosts - i/gallery/posts
func iGalleryPosts(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iGetWordMutedNotesCount - i/get-word-muted-notes-count
func iGetWordMutedNotesCount(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iNotifications - i/notifications
func iNotifications(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iPageLikes - i/page-likes
func iPageLikes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iPages - i/pages
func iPages(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iPin - i/pin
func iPin(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iReadAllMessagingMessages - i/read-all-messaging-messages
func iReadAllMessagingMessages(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iReadAllUnreadNotes - i/read-all-unread-notes
func iReadAllUnreadNotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iReadAnnouncement - i/read-announcement
func iReadAnnouncement(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iUnpin - i/unpin
func iUnpin(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iUpdate - i/update
func iUpdate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// iUserGroupInvites - i/user-group-invites
func iUserGroupInvites(c *gin.Context) {
	// TODO
	placeholder(c)
}

// muteCreate - mute/create
func muteCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// muteDelete - mute/delete
func muteDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// muteList - mute/list
func muteList(c *gin.Context) {
	// TODO
	placeholder(c)
}

// myApps - my/apps
func myApps(c *gin.Context) {
	// TODO
	placeholder(c)
}

// swRegister - sw/register
func swRegister(c *gin.Context) {
	// TODO
	placeholder(c)
}

// swUnregister - sw/unregister
func swUnregister(c *gin.Context) {
	// TODO
	placeholder(c)
}
