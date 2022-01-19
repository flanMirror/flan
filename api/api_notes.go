package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func channelsTimeline(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notes - notes
func notes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesChildren - notes/children
func notesChildren(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesConversation - notes/conversation
func notesConversation(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesCreate - notes/create
func notesCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesDelete - notes/delete
func notesDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesFavoritesCreate - notes/favorites/create
func notesFavoritesCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesFavoritesDelete - notes/favorites/delete
func notesFavoritesDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesFeatured - notes/featured
func notesFeatured(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesGlobalTimeline - notes/global-timeline
func notesGlobalTimeline(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesHybridTimeline - notes/hybrid-timeline
func notesHybridTimeline(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesLocalTimeline - notes/local-timeline
func notesLocalTimeline(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesMentions - notes/mentions
func notesMentions(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesPollsRecommendation - notes/polls/recommendation
func notesPollsRecommendation(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesPollsVote - notes/polls/vote
func notesPollsVote(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesReactions - notes/reactions
func notesReactions(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesRenotes - notes/renotes
func notesRenotes(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesReplies - notes/replies
func notesReplies(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesSearch - notes/search
func notesSearch(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesSearchByTag - notes/search-by-tag
func notesSearchByTag(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesShow - notes/show
func notesShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesState - notes/state
func notesState(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesThreadMutingCreate - notes/thread-muting/create
func notesThreadMutingCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesThreadMutingDelete - notes/thread-muting/delete
func notesThreadMutingDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesTimeline - notes/timeline
func notesTimeline(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesTranslate - notes/translate
func notesTranslate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesUnrenote - notes/unrenote
func notesUnrenote(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesUserListTimeline - notes/user-list-timeline
func notesUserListTimeline(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesWatchingCreate - notes/watching/create
func notesWatchingCreate(c *gin.Context) {
	// TODO
	placeholder(c)
}

// notesWatchingDelete - notes/watching/delete
func notesWatchingDelete(c *gin.Context) {
	// TODO
	placeholder(c)
}

// promoRead - promo/read
func promoRead(c *gin.Context) {
	// TODO
	placeholder(c)
}
