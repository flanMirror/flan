package payload

import (
	"github.com/volatiletech/null/v8"
	"random.chars.jp/git/misskey/api/response"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db/cache"
	"random.chars.jp/git/misskey/db/orm"
	"random.chars.jp/git/misskey/spec"
)

var (
	// Meta caches the response of meta
	// and is updated when
	// db.Meta is invalidated, using the populateMeta function
	Meta = data.New[response.Meta]()
	// MetaDetail caches the response of meta when detail is requested, with an unprivileged context
	// and is updated when
	// Meta is updated, using the populateMeta function
	MetaDetail = data.New[response.Meta]()
	// MetaAdmin caches the response of meta when detail is requested, with a privileged context
	// and is updated when
	// Meta is updated, using the populateMeta function
	MetaAdmin = data.New[response.MetaAdmin]()
)

func init() {
	cache.Ads.Register(func(value orm.AdSlice) {
		populateAds(value)
	})
	cache.Emojis.Register(func(value orm.EmojiSlice) {
		populateEmojis(value)
	})
	cache.LocalUserCount.Register(func(_ int64) {
		refreshMeta()
	})
	cache.ProxyAccount.Register(func(_ *orm.User) {
		refreshMeta()
	})
	cache.Meta.Register(func(metum *orm.Metum) {
		populateMeta(metum, response.Meta{
			MaintainerName:  metum.MaintainerName,
			MaintainerEmail: metum.MaintainerEmail,

			Version: spec.Target,

			Name:        metum.Name.String,
			URI:         config.URL,
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

func populateMeta(metum *orm.Metum, meta response.Meta) {
	Meta.Set(meta)

	meta.PinnedPages = metum.PinnedPages
	meta.PinnedClipID = &metum.PinnedClipId
	meta.CacheRemoteFiles = &metum.CacheRemoteFiles
	meta.ProxyRemoteFiles = &metum.ProxyRemoteFiles
	r := cache.LocalUserCount.Get() == 0
	meta.RequireSetup = &r

	// user packing is skipped here because it is not necessary
	// however, keep that in mind when catching up with upstream
	if p := cache.ProxyAccount.Get(); p != nil {
		proxyAccount := p
		meta.ProxyAccountName = &proxyAccount.Name
	} else {
		n := null.NewString("", false)
		meta.ProxyAccountName = &n
	}
	meta.Features = &response.MetaFeatures{
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

	metaAdmin := response.MetaAdmin{Meta: meta}
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

// refreshMeta calls populateMeta if available
func refreshMeta() {
	metum := cache.Meta.Get()
	if metum == nil {
		return
	}

	populateMeta(metum, Meta.Get())
}

func populateAds(ads orm.AdSlice) {
	metum := cache.Meta.Get()
	if metum == nil {
		return
	}

	payload := make([]response.MetaAd, len(ads))
	for i, ad := range ads {
		payload[i] = response.MetaAd{
			ID:       ad.ID,
			ImageURL: ad.ImageUrl,
			Place:    ad.Place,
			Ratio:    ad.Ratio,
			URL:      ad.URL,
		}
	}

	meta := Meta.Get()
	meta.Ads = payload
	populateMeta(metum, meta)
}

func populateEmojis(emojis orm.EmojiSlice) {
	metum := cache.Meta.Get()
	if metum == nil {
		return
	}

	payload := make([]response.MetaEmoji, len(emojis))
	for i, emoji := range emojis {
		payload[i] = response.MetaEmoji{
			ID:       emoji.ID,
			Aliases:  emoji.Aliases,
			Name:     emoji.Name,
			Category: emoji.Category,
			Host:     emoji.Host,
			URL:      emoji.URL,
		}
	}

	meta := Meta.Get()
	meta.Emojis = payload
	populateMeta(metum, meta)
}
