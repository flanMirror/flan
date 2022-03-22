package cache

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db/orm"
	"random.chars.jp/git/misskey/db/qhelper"
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
	// Ads caches ads as part of the meta response
	// and is invalidated when
	// TODO
	Ads = data.NewEager(func() (orm.AdSlice, bool) {
		if ads, err := orm.Ads(
			qhelper.ExpiresAt(time.Now()),
		).AllG(context.Background()); err != nil {
			log.Printf("error fetching ads: %s", err)
			return orm.AdSlice{}, false
		} else {
			return ads, true
		}
	})
	// Emojis caches emojis as part of the meta response
	// and is invalidated when
	// TODO
	Emojis = data.NewEager(func() (orm.EmojiSlice, bool) {
		if emojis, err := orm.Emojis(
			qmhelper.WhereIsNull("host"),
			qm.OrderBy("category asc, name asc"),
		).AllG(context.Background()); err != nil {
			log.Printf("error fetching emojis: %s", err)
			return orm.EmojiSlice{}, false
		} else {
			return emojis, true
		}
	})
	// Meta caches the instance metadata
	// and is invalidated when
	// TODO
	Meta = data.NewEager(func() (*orm.Metum, bool) {
		if m, err := fetchMeta(context.Background()); err != nil {
			log.Printf("error fetching meta: %s", err)
			return nil, false
		} else {
			return m, true
		}
	})
	// LocalUserCount caches the local user count
	// and is invalidated when
	// a local user is created or destroyed
	LocalUserCount = data.NewEager(func() (int64, bool) {
		if count, err := orm.Users(
			qmhelper.WhereIsNull("host"),
		).CountG(context.Background()); err != nil {
			log.Printf("error counting local users: %s", err)
			return 0, false
		} else {
			return count, true
		}
	})
	// ProxyAccount caches the proxy account
	// and is invalidated when
	// the proxy account is set or unset
	ProxyAccount = data.NewEager(func() (*orm.User, bool) {
		var meta *orm.Metum
		if m := Meta.Get(); m == nil {
			return nil, false
		} else {
			meta = m
		}

		if !meta.ProxyAccountId.Valid {
			return nil, false
		}

		if user, err := orm.FindUserG(context.Background(),
			meta.ProxyAccountId.String); err != nil {
			log.Printf("error fetching proxy account: %s", err)
			return nil, false
		} else {
			return user, true
		}
	})

	metaLock = sync.RWMutex{}
)

const (
	MaxNoteTextLength     = 8192
	MaxImageCommentLength = 512
)

func fetchMeta(ctx context.Context) (*orm.Metum, error) {
	// upstream does this because apparently some "bugs" from the past
	// has caused multiple entries to appear
	// reading the code it's clearly caused by a race condition which we've
	// prevented here, however since we'll be reading from legacy databases
	// we still need to keep this insanity around
	metaLock.RLock()
	defer metaLock.RUnlock()

	if m, err := orm.Meta(
		qm.OrderBy("id desc"),
	).OneG(ctx); err != nil {
		return nil, err
	} else {
		if m == nil {
			metaLock.RUnlock()
			defer metaLock.RLock()

			metaLock.Lock()
			defer metaLock.Unlock()
			log.Print("got nil meta from database, setting defaults")
			m = &orm.Metum{ID: "x"}
			if err = orm.Meta().BindG(ctx, m); err != nil {
				return nil, err
			}
		}
		if m.MaxNoteTextLength > MaxNoteTextLength {
			m.MaxNoteTextLength = MaxNoteTextLength
		}
		return m, nil
	}
}
