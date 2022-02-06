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

	meta.UseStarForReactionFallback = &metum.UseStarForReactionFallback
	// TODO: need to test if this becomes nil after authentication is ready
	meta.PinnedUsers = metum.PinnedUsers
	meta.HiddenTags = metum.HiddenTags
	meta.BlockedHosts = metum.BlockedHosts
	meta.HCaptchaSecretKey = metum.HcaptchaSecretKey.Ptr()
	meta.RecaptchaSecretKey = metum.RecaptchaSecretKey.Ptr()
	meta.ProxyAccountID = metum.ProxyAccountId.Ptr()
	meta.TwitterConsumerKey = metum.TwitterConsumerKey.Ptr()
	meta.TwitterConsumerSecret = metum.TwitterConsumerSecret.Ptr()
	meta.GithubClientID = metum.GithubClientId.Ptr()
	meta.GithubClientSecret = metum.GithubClientSecret.Ptr()
	meta.DiscordClientID = metum.DiscordClientId.Ptr()
	meta.DiscordClientSecret = metum.DiscordClientSecret.Ptr()
	meta.SummalyProxy = metum.SummalyProxy.Ptr()
	meta.Email = metum.Email.Ptr()
	meta.SmtpSecure = &metum.SmtpSecure
	meta.SmtpHost = metum.SmtpHost.Ptr()
	meta.SmtpPort = metum.SmtpPort.Ptr()
	meta.SmtpUser = metum.SmtpUser.Ptr()
	meta.SmtpPass = metum.SmtpPass.Ptr()
	meta.SwPrivateKey = metum.SwPrivateKey.Ptr()
	meta.UseObjectStorage = &metum.UseObjectStorage
	meta.ObjectStorageBaseURL = metum.ObjectStorageBaseUrl.Ptr()
	meta.ObjectStorageBucket = metum.ObjectStorageBucket.Ptr()
	meta.ObjectStoragePrefix = metum.ObjectStoragePrefix.Ptr()
	meta.ObjectStorageEndpoint = metum.ObjectStorageEndpoint.Ptr()
	meta.ObjectStorageRegion = metum.ObjectStorageRegion.Ptr()
	meta.ObjectStoragePort = metum.ObjectStoragePort.Ptr()
	meta.ObjectStorageAccessKey = metum.ObjectStorageAccessKey.Ptr()
	meta.ObjectStorageSecretKey = metum.ObjectStorageSecretKey.Ptr()
	meta.ObjectStorageUseSSL = &metum.ObjectStorageUseSSL
	meta.ObjectStorageUseProxy = &metum.ObjectStorageUseProxy
	meta.ObjectStorageSetPublicRead = &metum.ObjectStorageSetPublicRead
	meta.ObjectStorageS3ForcePathStyle = &metum.ObjectStorageS3ForcePathStyle
	meta.DeeplAuthKey = metum.DeeplAuthKey.Ptr()
	meta.DeeplIsPro = &metum.DeeplIsPro
	MetaAdmin.Set(meta)
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
