package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"log"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db/models"
	"sync"
	"time"
)

/* this is ridiculous, why the hell is ads and emoji sent with everything else??
this is incredibly inefficient and the upstream isn't even cached

in here I have to use some insanely ugly invalidation but at least it's still
going to be way faster than upstream

the init invalidation goes like this: first LocalUserCount is invalidated and its
invalidation handler returns prematurely due to having a nil Meta cache. Meta is
then invalidated which populates the main Meta cache in payload caches, after that
Ads, Emojis and ProxyAccount are subsequently invalidated, they populate those
three fields and calls the populateMeta function individually to update the other
meta variant structs. after this insane initial pass is done, invalidation to any
of these caches would work at any point onwards */

var (
	// Ads caches database entity models.AdSlice
	Ads = data.NewEager(func() interface{} {
		if ads, err := models.Ads(
			qm.Where(
				"\"expiresAt\" >= (? at time zone 'utc')",
				time.Now().UTC().Format(TimestampFormat),
			),
		).AllG(context.Background()); err != nil {
			log.Printf("error fetching ads: %s", err)
			return nil
		} else {
			return ads
		}
	})
	// Emojis caches database entity models.EmojiSlice
	Emojis = data.NewEager(func() interface{} {
		if emojis, err := models.Emojis(
			qmhelper.WhereIsNull("host"),
			qm.OrderBy("category asc, name asc"),
		).AllG(context.Background()); err != nil {
			log.Printf("error fetching emojis: %s", err)
			return nil
		} else {
			return emojis
		}
	})
	// Meta caches database entity *models.Metum
	Meta = data.NewEager(func() interface{} {
		if m, err := fetchMeta(context.Background()); err != nil {
			log.Printf("error fetching meta: %s", err)
			return nil
		} else {
			return m
		}
	})
	// LocalUserCount caches int64 amount of local users
	LocalUserCount = data.NewEager(func() interface{} {
		if count, err := models.Users(
			qmhelper.WhereIsNull("host"),
		).CountG(context.Background()); err != nil {
			log.Printf("error counting local users: %s", err)
			return nil
		} else {
			return count
		}
	})
	// ProxyAccount caches database entity *models.User proxy account
	ProxyAccount = data.NewEager(func() interface{} {
		var meta *models.Metum
		if m := Meta.Get(); m == nil {
			return nil
		} else {
			meta = m.(*models.Metum)
		}

		if !meta.ProxyAccountId.Valid {
			return nil
		}

		if user, err := models.Users(
			qm.Where("id = ?", meta.ProxyAccountId.String),
		).OneG(context.Background()); err != nil {
			log.Printf("error fetching proxy account: %s", err)
			return nil
		} else {
			return user
		}
	})

	metaLock = sync.RWMutex{}
)

const (
	MaxNoteTextLength     = 8192
	MaxImageCommentLength = 512
)

func fetchMeta(ctx context.Context) (*models.Metum, error) {
	// upstream does this because apparently some "bugs" from the past
	// has caused multiple entries to appear
	// reading the code it's clearly caused by a race condition which we've
	// prevented here, however since we'll be reading from legacy databases
	// we still need to keep this insanity around
	metaLock.RLock()
	defer metaLock.RUnlock()

	if m, err := models.Meta(qm.OrderBy("id desc")).OneG(ctx); err != nil {
		return nil, err
	} else {
		if m == nil {
			metaLock.RUnlock()
			defer metaLock.RLock()

			metaLock.Lock()
			defer metaLock.Unlock()
			log.Print("got nil meta from database, setting defaults")
			m = &models.Metum{ID: "x"}
			if err = models.Meta().BindG(ctx, m); err != nil {
				return nil, err
			}
		}
		if m.MaxNoteTextLength > MaxNoteTextLength {
			m.MaxNoteTextLength = MaxNoteTextLength
		}
		return m, nil
	}
}
