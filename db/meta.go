package db

import (
	"context"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"random.chars.jp/git/misskey/db/models"
)

var metaCache *models.Metum

func FetchMeta(ctx context.Context) (*models.Metum, error) {
	if metaCache != nil {
		return metaCache, nil
	}

	// don't blame me for this insanity, this is what upstream does, and we are just a drop-in replacement
	if m, err := fetchMeta(ctx); err != nil {
		return nil, err
	} else {
		if m == nil {
			log.Print("got nil meta from database, setting defaults")
			m = &models.Metum{ID: "x"}
			if err = models.Meta().BindG(ctx, m); err != nil {
				return nil, err
			}
		}
		metaCache = m
		return m, nil
	}
}

func InvalidateMeta() {
	metaCache = nil
}

func fetchMeta(ctx context.Context) (*models.Metum, error) {
	return models.Meta(qm.OrderBy("id desc")).OneG(ctx)
}
