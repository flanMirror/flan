package openapi

import (
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
func gamesReversiGames(ctx Context) {
	// TODO
	placeholder(ctx)
}

// gamesReversiGamesShow - games/reversi/games/show
func gamesReversiGamesShow(ctx Context) {
	// TODO
	placeholder(ctx)
}

// gamesReversiGamesSurrender - games/reversi/games/surrender
func gamesReversiGamesSurrender(ctx Context) {
	// TODO
	placeholder(ctx)
}

// gamesReversiInvitations - games/reversi/invitations
func gamesReversiInvitations(ctx Context) {
	// TODO
	placeholder(ctx)
}

// gamesReversiMatch - games/reversi/match
func gamesReversiMatch(ctx Context) {
	// TODO
	placeholder(ctx)
}

// gamesReversiMatchCancel - games/reversi/match/cancel
func gamesReversiMatchCancel(ctx Context) {
	// TODO
	placeholder(ctx)
}
