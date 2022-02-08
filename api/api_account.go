package openapi

import (
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
func blockingCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// blockingDelete - blocking/delete
func blockingDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// blockingList - blocking/list
func blockingList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// clipsAddNote - clips/add-note
func clipsAddNote(ctx Context) {
	// TODO
	placeholder(ctx)
}

// clipsNotes - clips/notes
func clipsNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// i - i
func i(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iFavorites - i/favorites
func iFavorites(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iGalleryLikes - i/gallery/likes
func iGalleryLikes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iGalleryPosts - i/gallery/posts
func iGalleryPosts(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iGetWordMutedNotesCount - i/get-word-muted-notes-count
func iGetWordMutedNotesCount(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iNotifications - i/notifications
func iNotifications(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iPageLikes - i/page-likes
func iPageLikes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iPages - i/pages
func iPages(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iPin - i/pin
func iPin(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iReadAllMessagingMessages - i/read-all-messaging-messages
func iReadAllMessagingMessages(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iReadAllUnreadNotes - i/read-all-unread-notes
func iReadAllUnreadNotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iReadAnnouncement - i/read-announcement
func iReadAnnouncement(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iUnpin - i/unpin
func iUnpin(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iUpdate - i/update
func iUpdate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// iUserGroupInvites - i/user-group-invites
func iUserGroupInvites(ctx Context) {
	// TODO
	placeholder(ctx)
}

// muteCreate - mute/create
func muteCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// muteDelete - mute/delete
func muteDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// muteList - mute/list
func muteList(ctx Context) {
	// TODO
	placeholder(ctx)
}

// myApps - my/apps
func myApps(ctx Context) {
	// TODO
	placeholder(ctx)
}

// swRegister - sw/register
func swRegister(ctx Context) {
	// TODO
	placeholder(ctx)
}

// swUnregister - sw/unregister
func swUnregister(ctx Context) {
	// TODO
	placeholder(ctx)
}
