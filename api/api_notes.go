package api

import (
	"net/http"
)

func init() {
	route{http.MethodPost, "channels/timeline", channelsTimeline}.register()
	route{http.MethodPost, "notes", notes}.register()
	route{http.MethodPost, "notes/children", notesChildren}.register()
	route{http.MethodPost, "notes/conversation", notesConversation}.register()
	route{http.MethodPost, "notes/create", notesCreate}.register()
	route{http.MethodPost, "notes/delete", notesDelete}.register()
	route{http.MethodPost, "notes/favorites/create", notesFavoritesCreate}.register()
	route{http.MethodPost, "notes/favorites/delete", notesFavoritesDelete}.register()
	route{http.MethodPost, "notes/featured", notesFeatured}.register()
	route{http.MethodPost, "notes/global-timeline", notesGlobalTimeline}.register()
	route{http.MethodPost, "notes/hybrid-timeline", notesHybridTimeline}.register()
	route{http.MethodPost, "notes/local-timeline", notesLocalTimeline}.register()
	route{http.MethodPost, "notes/mentions", notesMentions}.register()
	route{http.MethodPost, "notes/polls/recommendation", notesPollsRecommendation}.register()
	route{http.MethodPost, "notes/polls/vote", notesPollsVote}.register()
	route{http.MethodPost, "notes/reactions", notesReactions}.register()
	route{http.MethodPost, "notes/renotes", notesRenotes}.register()
	route{http.MethodPost, "notes/replies", notesReplies}.register()
	route{http.MethodPost, "notes/search", notesSearch}.register()
	route{http.MethodPost, "notes/search-by-tag", notesSearchByTag}.register()
	route{http.MethodPost, "notes/show", notesShow}.register()
	route{http.MethodPost, "notes/state", notesState}.register()
	route{http.MethodPost, "notes/thread-muting/create", notesThreadMutingCreate}.register()
	route{http.MethodPost, "notes/thread-muting/delete", notesThreadMutingDelete}.register()
	route{http.MethodPost, "notes/timeline", notesTimeline}.register()
	route{http.MethodPost, "notes/translate", notesTranslate}.register()
	route{http.MethodPost, "notes/unrenote", notesUnrenote}.register()
	route{http.MethodPost, "notes/user-list-timeline", notesUserListTimeline}.register()
	route{http.MethodPost, "notes/watching/create", notesWatchingCreate}.register()
	route{http.MethodPost, "notes/watching/delete", notesWatchingDelete}.register()
	route{http.MethodPost, "promo/read", promoRead}.register()
}

// channelsTimeline - channels/timeline
func channelsTimeline(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notes - notes
func notes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesChildren - notes/children
func notesChildren(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesConversation - notes/conversation
func notesConversation(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesCreate - notes/create
func notesCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesDelete - notes/delete
func notesDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesFavoritesCreate - notes/favorites/create
func notesFavoritesCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesFavoritesDelete - notes/favorites/delete
func notesFavoritesDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesFeatured - notes/featured
func notesFeatured(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesGlobalTimeline - notes/global-timeline
func notesGlobalTimeline(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesHybridTimeline - notes/hybrid-timeline
func notesHybridTimeline(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesLocalTimeline - notes/local-timeline
func notesLocalTimeline(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesMentions - notes/mentions
func notesMentions(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesPollsRecommendation - notes/polls/recommendation
func notesPollsRecommendation(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesPollsVote - notes/polls/vote
func notesPollsVote(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesReactions - notes/reactions
func notesReactions(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesRenotes - notes/renotes
func notesRenotes(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesReplies - notes/replies
func notesReplies(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesSearch - notes/search
func notesSearch(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesSearchByTag - notes/search-by-tag
func notesSearchByTag(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesShow - notes/show
func notesShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesState - notes/state
func notesState(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesThreadMutingCreate - notes/thread-muting/create
func notesThreadMutingCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesThreadMutingDelete - notes/thread-muting/delete
func notesThreadMutingDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesTimeline - notes/timeline
func notesTimeline(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesTranslate - notes/translate
func notesTranslate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesUnrenote - notes/unrenote
func notesUnrenote(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesUserListTimeline - notes/user-list-timeline
func notesUserListTimeline(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesWatchingCreate - notes/watching/create
func notesWatchingCreate(ctx Context) {
	// TODO
	placeholder(ctx)
}

// notesWatchingDelete - notes/watching/delete
func notesWatchingDelete(ctx Context) {
	// TODO
	placeholder(ctx)
}

// promoRead - promo/read
func promoRead(ctx Context) {
	// TODO
	placeholder(ctx)
}
