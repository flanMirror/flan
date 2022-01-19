package openapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	route{http.MethodPost, "games/reversi/games", gamesReversiGames}.register()
	route{http.MethodPost, "games/reversi/games/show", gamesReversiGamesShow}.register()
	route{http.MethodPost, "games/reversi/games/surrender", gamesReversiGamesSurrender}.register()
	route{http.MethodPost, "games/reversi/invitations", gamesReversiInvitations}.register()
	route{http.MethodPost, "games/reversi/match", gamesReversiMatch}.register()
	route{http.MethodPost, "games/reversi/match/cancel", gamesReversiMatchCancel}.register()
}

// gamesReversiGames - games/reversi/games
func gamesReversiGames(c *gin.Context) {
	// TODO
	placeholder(c)
}

// gamesReversiGamesShow - games/reversi/games/show
func gamesReversiGamesShow(c *gin.Context) {
	// TODO
	placeholder(c)
}

// gamesReversiGamesSurrender - games/reversi/games/surrender
func gamesReversiGamesSurrender(c *gin.Context) {
	// TODO
	placeholder(c)
}

// gamesReversiInvitations - games/reversi/invitations
func gamesReversiInvitations(c *gin.Context) {
	// TODO
	placeholder(c)
}

// gamesReversiMatch - games/reversi/match
func gamesReversiMatch(c *gin.Context) {
	// TODO
	placeholder(c)
}

// gamesReversiMatchCancel - games/reversi/match/cancel
func gamesReversiMatchCancel(c *gin.Context) {
	// TODO
	placeholder(c)
}
