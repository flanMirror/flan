package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "blocking/create", blockingCreate}.register()
	route{http.MethodPost, "blocking/delete", blockingDelete}.register()
	route{http.MethodPost, "blocking/list", blockingList}.register()
	route{http.MethodPost, "clips/add-note", clipsAddNote}.register()
	route{http.MethodPost, "clips/notes", clipsNotes}.register()
	route{http.MethodPost, "i", i}.register()
	route{http.MethodPost, "i/favorites", iFavorites}.register()
	route{http.MethodPost, "i/gallery/likes", iGalleryLikes}.register()
	route{http.MethodPost, "i/gallery/posts", iGalleryPosts}.register()
	route{http.MethodPost, "i/get-word-muted-notes-count", iGetWordMutedNotesCount}.register()
	route{http.MethodPost, "i/notifications", iNotifications}.register()
	route{http.MethodPost, "i/page-likes", iPageLikes}.register()
	route{http.MethodPost, "i/pages", iPages}.register()
	route{http.MethodPost, "i/pin", iPin}.register()
	route{http.MethodPost, "i/read-all-messaging-messages", iReadAllMessagingMessages}.register()
	route{http.MethodPost, "i/read-all-unread-notes", iReadAllUnreadNotes}.register()
	route{http.MethodPost, "i/read-announcement", iReadAnnouncement}.register()
	route{http.MethodPost, "i/unpin", iUnpin}.register()
	route{http.MethodPost, "i/update", iUpdate}.register()
	route{http.MethodPost, "i/user-group-invites", iUserGroupInvites}.register()
	route{http.MethodPost, "mute/create", muteCreate}.register()
	route{http.MethodPost, "mute/delete", muteDelete}.register()
	route{http.MethodPost, "mute/list", muteList}.register()
	route{http.MethodPost, "my/apps", myApps}.register()
	route{http.MethodPost, "sw/register", swRegister}.register()
	route{http.MethodPost, "sw/unregister", swUnregister}.register()
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
