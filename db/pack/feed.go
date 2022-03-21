//go:build !bugforbug

package pack

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/orm"
	"random.chars.jp/git/misskey/feed"
)

var generator = "Misskey"

var ErrNoProfile = errors.New("no profile found")

func Feed(ctx context.Context, user *orm.User) (*feed.Emitter, error) {
	userRef := "@" + user.Username
	if user.Host.Valid {
		userRef += "@" + user.Host.String
	}
	link := config.URL + "/" + userRef
	name := user.Username
	if user.Name.Valid {
		name = user.Name.String
	}
	author := feed.Author{
		Link: &link,
		Name: &name,
	}

	var profile *orm.UserProfile
	if p, err := orm.FindUserProfileG(ctx, user.ID); err != nil {
		return nil, err
	} else {
		if p == nil {
			return nil, ErrNoProfile
		}
		profile = p
	}

	var notes orm.NoteSlice
	if n, err := orm.Notes(
		qm.Where(`"userId" = ?`, user.ID),
		qm.Where(`"renoteId" is null`),
		qm.Where(`visibility in ('public', 'home')`),

		qm.OrderBy(`"createdAt" desc`),

		qm.Limit(20),
	).AllG(ctx); err != nil {
		return nil, err
	} else {
		notes = n
	}

	var updated time.Time
	if len(notes) > 0 {
		updated = notes[0].CreatedAt
	}

	public := profile.FfVisibility == "public"
	var (
		followingCount string
		followersCount string
		description    string
	)
	if public {
		followingCount = strconv.Itoa(user.FollowingCount)
		followersCount = strconv.Itoa(user.FollowersCount)
	} else {
		followingCount = "?"
		followersCount = "?"
	}
	if profile.Description.Valid {
		description = " · " + profile.Description.String
	}

	emitter := feed.New(feed.Options{
		ID:        link,
		Title:     name + " (" + userRef + ")",
		Updated:   &updated,
		Generator: &generator,
		Description: s(
			strconv.Itoa(user.NotesCount) + " Notes, " +
				followingCount + " Following, " +
				followersCount + " Followers" +
				description),
		Link:  author.Link,
		Image: user.AvatarUrl.Ptr(),
		FeedLinks: map[string]string{
			"json": link + ".json",
			"atom": link + ".atom",
		},
		Author:    &author,
		Copyright: &name,
	})

	emitter.Items = make([]feed.Item, len(notes))
	for i, note := range notes {
		var files orm.DriveFileSlice
		if len(note.FileIds) > 0 {
			if f, err := orm.DriveFiles(qm.Where("id in ?", note.FileIds)).AllG(ctx); err != nil {
				return nil, err
			} else {
				files = f
			}
		}
		var url *string
		for _, file := range files {
			if strings.HasPrefix(file.Type, "image/") {
				// an empty string is returned to represent nil
				if str := db.GetPublicURL(file, false, nil); str != "" {
					url = &str
				}
				break
			}
		}
		d := note.CreatedAt.UTC()
		emitter.Items[i] = feed.Item{
			Title:       "New note by " + name,
			Link:        config.URL + "/notes/" + note.ID,
			Date:        &d,
			Description: note.CW.Ptr(),
			Content:     note.Text.Ptr(),
			Image:       url,
		}
	}

	return emitter, nil
}

func s(str string) *string {
	return &str
}
