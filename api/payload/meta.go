package payload

import (
	"github.com/volatiletech/null/v8"
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/models"
	"random.chars.jp/git/misskey/spec"
)

var (
	// Meta caches structs.Meta.
	Meta = data.New()
	// MetaDetail caches structs.Meta with detail fields populated.
	MetaDetail = data.New()
	// MetaAdmin caches structs.Meta with detail and admin-only fields populated.
	MetaAdmin = data.New()
)

func init() {
	db.Ads.Register(func(data interface{}) {
		populateAds(data.(models.AdSlice))
	})
	db.Emojis.Register(func(data interface{}) {
		populateEmojis(data.(models.EmojiSlice))
	})
	db.LocalUserCount.Register(func(data interface{}) {
		refreshMeta()
	})
	db.ProxyAccount.Register(func(data interface{}) {
		refreshMeta()
	})
	db.Meta.Register(func(data interface{}) {
		metum := data.(*models.Metum)

		populateMeta(metum, structs.Meta{
			MaintainerName:  metum.MaintainerName,
			MaintainerEmail: metum.MaintainerEmail,

			Version: spec.Target,

			Name:        metum.Name.String,
			URI:         config.Web.URL,
			Description: metum.Description,
			Langs:       metum.Langs,
			TosURL:      metum.ToSUrl,
			//RepositoryURL: metum.RepositoryUrl,
			RepositoryURL: spec.Repository,
			FeedbackURL:   metum.FeedbackUrl.String,

			Secure: config.HTTPS,

			DisableRegistration:          metum.DisableRegistration,
			DisableLocalTimeline:         metum.DisableLocalTimeline,
			DisableGlobalTimeline:        metum.DisableGlobalTimeline,
			DriveCapacityPerLocalUserMb:  metum.LocalDriveCapacityMb,
			DriveCapacityPerRemoteUserMb: metum.RemoteDriveCapacityMb,
			EmailRequiredForSignup:       metum.EmailRequiredForSignup,
			EnableHcaptcha:               metum.EnableHcaptcha,
			HcaptchaSiteKey:              metum.HcaptchaSiteKey,
			EnableRecaptcha:              metum.EnableRecaptcha,
			RecaptchaSiteKey:             metum.RecaptchaSiteKey,
			SwPublickey:                  metum.SwPublicKey,
			MascotImageURL:               metum.MascotImageUrl.String,
			BannerURL:                    metum.BannerUrl.String,
			ErrorImageURL:                metum.ErrorImageUrl.String,
			IconURL:                      metum.IconUrl,
			BackgroundImageURL:           metum.BackgroundImageUrl,
			LogoImageURL:                 metum.LogoImageUrl,
			MaxNoteTextLength:            metum.MaxNoteTextLength,
			// emoji and ads are cached and populated separately
			EnableEmail: metum.EnableEmail,

			EnableTwitterIntegration: metum.EnableTwitterIntegration,
			EnableGithubIntegration:  metum.EnableGithubIntegration,
			EnableDiscordIntegration: metum.EnableDiscordIntegration,

			EnableServiceWorker: metum.EnableServiceWorker,

			TranslatorAvailable: metum.DeeplAuthKey.Valid,
		})
	})
}

func populateMeta(metum *models.Metum, meta structs.Meta) {
	Meta.Set(meta)

	meta.PinnedPages = metum.PinnedPages
	meta.PinnedClipID = &metum.PinnedClipId
	meta.CacheRemoteFiles = &metum.CacheRemoteFiles
	meta.ProxyRemoteFiles = &metum.ProxyRemoteFiles
	if luc, ok := db.LocalUserCount.Get().(int64); ok {
		r := luc == 0
		meta.RequireSetup = &r
	}
	// user packing is skipped here because it is not necessary
	// however, keep that in mind when catching up with upstream
	if p := db.ProxyAccount.Get(); p != nil {
		proxyAccount := p.(*models.User)
		meta.ProxyAccountName = &proxyAccount.Name
	} else {
		n := null.NewString("", false)
		meta.ProxyAccountName = &n
	}
	meta.Features = &structs.MetaFeatures{
		Registration:           !metum.DisableRegistration,
		LocalTimeLine:          !metum.DisableLocalTimeline,
		GlobalTimeLine:         !metum.DisableGlobalTimeline,
		EmailRequiredForSignup: metum.EmailRequiredForSignup,
		Elasticsearch:          config.External.Elasticsearch.Enable,
		HCaptcha:               metum.EnableHcaptcha,
		Recaptcha:              metum.EnableRecaptcha,
		ObjectStorage:          metum.UseObjectStorage,
		Twitter:                metum.EnableTwitterIntegration,
		Github:                 metum.EnableGithubIntegration,
		Discord:                metum.EnableDiscordIntegration,
		ServiceWorker:          metum.EnableServiceWorker,
		MiAuth:                 true,
	}
	MetaDetail.Set(meta)

	metaAdmin := structs.MetaAdmin{Meta: meta}
	metaAdmin.UseStarForReactionFallback = metum.UseStarForReactionFallback
	metaAdmin.PinnedUsers = metum.PinnedUsers
	metaAdmin.HiddenTags = metum.HiddenTags
	metaAdmin.BlockedHosts = metum.BlockedHosts
	metaAdmin.HCaptchaSecretKey = metum.HcaptchaSecretKey
	metaAdmin.RecaptchaSecretKey = metum.RecaptchaSecretKey
	metaAdmin.ProxyAccountID = metum.ProxyAccountId
	metaAdmin.TwitterConsumerKey = metum.TwitterConsumerKey
	metaAdmin.TwitterConsumerSecret = metum.TwitterConsumerSecret
	metaAdmin.GithubClientID = metum.GithubClientId
	metaAdmin.GithubClientSecret = metum.GithubClientSecret
	metaAdmin.DiscordClientID = metum.DiscordClientId
	metaAdmin.DiscordClientSecret = metum.DiscordClientSecret
	metaAdmin.SummalyProxy = metum.SummalyProxy
	metaAdmin.Email = metum.Email
	metaAdmin.SmtpSecure = metum.SmtpSecure
	metaAdmin.SmtpHost = metum.SmtpHost
	metaAdmin.SmtpPort = metum.SmtpPort
	metaAdmin.SmtpUser = metum.SmtpUser
	metaAdmin.SmtpPass = metum.SmtpPass
	metaAdmin.SwPrivateKey = metum.SwPrivateKey
	metaAdmin.UseObjectStorage = metum.UseObjectStorage
	metaAdmin.ObjectStorageBaseURL = metum.ObjectStorageBaseUrl
	metaAdmin.ObjectStorageBucket = metum.ObjectStorageBucket
	metaAdmin.ObjectStoragePrefix = metum.ObjectStoragePrefix
	metaAdmin.ObjectStorageEndpoint = metum.ObjectStorageEndpoint
	metaAdmin.ObjectStorageRegion = metum.ObjectStorageRegion
	metaAdmin.ObjectStoragePort = metum.ObjectStoragePort
	metaAdmin.ObjectStorageAccessKey = metum.ObjectStorageAccessKey
	metaAdmin.ObjectStorageSecretKey = metum.ObjectStorageSecretKey
	metaAdmin.ObjectStorageUseSSL = metum.ObjectStorageUseSSL
	metaAdmin.ObjectStorageUseProxy = metum.ObjectStorageUseProxy
	metaAdmin.ObjectStorageSetPublicRead = metum.ObjectStorageSetPublicRead
	metaAdmin.ObjectStorageS3ForcePathStyle = metum.ObjectStorageS3ForcePathStyle
	metaAdmin.DeeplAuthKey = metum.DeeplAuthKey
	metaAdmin.DeeplIsPro = metum.DeeplIsPro
	MetaAdmin.Set(metaAdmin)
}

func getMetum() *models.Metum {
	if m := db.Meta.Get(); m == nil {
		return nil
	} else {
		return m.(*models.Metum)
	}
}

// refreshMeta calls populateMeta if available
func refreshMeta() {
	metum := getMetum()
	if metum == nil {
		return
	}

	if m := Meta.Get(); m == nil {
		return
	} else {
		populateMeta(metum, m.(structs.Meta))
	}
}

func populateAds(ads models.AdSlice) {
	metum := getMetum()
	if metum == nil {
		return
	}

	if m := Meta.Get(); m == nil {
		return
	} else {
		payload := make([]structs.MetaAd, len(ads))
		for i, ad := range ads {
			payload[i] = structs.MetaAd{
				ID:       ad.ID,
				ImageURL: ad.ImageUrl,
				Place:    ad.Place,
				Ratio:    ad.Ratio,
				URL:      ad.URL,
			}
		}

		meta := m.(structs.Meta)
		meta.Ads = payload
		populateMeta(metum, meta)
	}
}

func populateEmojis(emojis models.EmojiSlice) {
	metum := getMetum()
	if metum == nil {
		return
	}

	if m := Meta.Get(); m == nil {
		return
	} else {
		payload := make([]structs.MetaEmoji, len(emojis))
		for i, emoji := range emojis {
			payload[i] = structs.MetaEmoji{
				ID:       emoji.ID,
				Aliases:  emoji.Aliases,
				Name:     emoji.Name,
				Category: emoji.Category,
				Host:     emoji.Host,
				URL:      emoji.URL,
			}
		}

		meta := m.(structs.Meta)
		meta.Emojis = payload
		populateMeta(metum, meta)
	}
}
