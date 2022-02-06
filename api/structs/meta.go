package structs

import "github.com/volatiletech/null/v8"

type Meta struct {
	MaintainerName                null.String   `json:"maintainerName"`
	MaintainerEmail               null.String   `json:"maintainerEmail"`
	Version                       string        `json:"version"`
	Name                          string        `json:"name"`
	URI                           string        `json:"uri"`
	Description                   null.String   `json:"description"`
	Langs                         []string      `json:"langs"`
	TosURL                        null.String   `json:"tosUrl"`
	RepositoryURL                 string        `json:"repositoryUrl"`
	FeedbackURL                   string        `json:"feedbackUrl"`
	Secure                        bool          `json:"secure"`
	DisableRegistration           bool          `json:"disableRegistration"`
	DisableLocalTimeline          bool          `json:"disableLocalTimeline"`
	DisableGlobalTimeline         bool          `json:"disableGlobalTimeline"`
	DriveCapacityPerLocalUserMb   int           `json:"driveCapacityPerLocalUserMb"`
	DriveCapacityPerRemoteUserMb  int           `json:"driveCapacityPerRemoteUserMb"`
	CacheRemoteFiles              *bool         `json:"cacheRemoteFiles,omitempty"`
	ProxyRemoteFiles              *bool         `json:"proxyRemoteFiles,omitempty"`
	EmailRequiredForSignup        bool          `json:"emailRequiredForSignup"`
	EnableHcaptcha                bool          `json:"enableHcaptcha"`
	HcaptchaSiteKey               null.String   `json:"hcaptchaSiteKey"`
	EnableRecaptcha               bool          `json:"enableRecaptcha"`
	RecaptchaSiteKey              null.String   `json:"recaptchaSiteKey"`
	SwPublickey                   null.String   `json:"swPublickey"`
	MascotImageURL                string        `json:"mascotImageUrl"`
	BannerURL                     string        `json:"bannerUrl"`
	ErrorImageURL                 string        `json:"errorImageUrl"`
	IconURL                       null.String   `json:"iconUrl"`
	BackgroundImageURL            null.String   `json:"backgroundImageUrl"`
	LogoImageURL                  null.String   `json:"logoImageUrl"`
	MaxNoteTextLength             int           `json:"maxNoteTextLength"`
	Emojis                        []MetaEmoji   `json:"emojis"`
	Ads                           []MetaAd      `json:"ads"`
	RequireSetup                  *bool         `json:"requireSetup,omitempty"`
	EnableEmail                   bool          `json:"enableEmail"`
	EnableTwitterIntegration      bool          `json:"enableTwitterIntegration"`
	EnableGithubIntegration       bool          `json:"enableGithubIntegration"`
	EnableDiscordIntegration      bool          `json:"enableDiscordIntegration"`
	EnableServiceWorker           bool          `json:"enableServiceWorker"`
	TranslatorAvailable           bool          `json:"translatorAvailable"`
	ProxyAccountName              *null.String  `json:"proxyAccountName,omitempty"`
	Features                      *MetaFeatures `json:"features,omitempty"`
	UseStarForReactionFallback    *bool         `json:"useStarForReactionFallback,omitempty"`
	PinnedUsers                   []string      `json:"pinnedUsers,omitempty"`
	PinnedPages                   []string      `json:"pinnedPages,omitempty"`
	PinnedClipID                  *null.String  `json:"pinnedClipId,omitempty"`
	HiddenTags                    []string      `json:"hiddenTags,omitempty"`
	BlockedHosts                  []string      `json:"blockedHosts,omitempty"`
	HCaptchaSecretKey             *string       `json:"hcaptchaSecretKey,omitempty"`
	RecaptchaSecretKey            *string       `json:"recaptchaSecretKey,omitempty"`
	ProxyAccountID                *string       `json:"proxyAccountId,omitempty"`
	TwitterConsumerKey            *string       `json:"twitterConsumerKey,omitempty"`
	TwitterConsumerSecret         *string       `json:"twitterConsumerSecret,omitempty"`
	GithubClientID                *string       `json:"githubClientId,omitempty"`
	GithubClientSecret            *string       `json:"githubClientSecret,omitempty"`
	DiscordClientID               *string       `json:"discordClientId,omitempty"`
	DiscordClientSecret           *string       `json:"discordClientSecret,omitempty"`
	SummalyProxy                  *string       `json:"summalyProxy,omitempty"`
	Email                         *string       `json:"email,omitempty"`
	SmtpSecure                    *bool         `json:"smtpSecure,omitempty"`
	SmtpHost                      *string       `json:"smtpHost,omitempty"`
	SmtpPort                      *int          `json:"smtpPort,omitempty"`
	SmtpUser                      *string       `json:"smtpUser,omitempty"`
	SmtpPass                      *string       `json:"smtpPass,omitempty"`
	SwPrivateKey                  *string       `json:"swPrivateKey,omitempty"`
	UseObjectStorage              *bool         `json:"useObjectStorage,omitempty"`
	ObjectStorageBaseURL          *string       `json:"objectStorageBaseUrl,omitempty"`
	ObjectStorageBucket           *string       `json:"objectStorageBucket,omitempty"`
	ObjectStoragePrefix           *string       `json:"objectStoragePrefix,omitempty"`
	ObjectStorageEndpoint         *string       `json:"objectStorageEndpoint,omitempty"`
	ObjectStorageRegion           *string       `json:"objectStorageRegion,omitempty"`
	ObjectStoragePort             *int          `json:"objectStoragePort,omitempty"`
	ObjectStorageAccessKey        *string       `json:"objectStorageAccessKey,omitempty"`
	ObjectStorageSecretKey        *string       `json:"objectStorageSecretKey,omitempty"`
	ObjectStorageUseSSL           *bool         `json:"objectStorageUseSSL,omitempty"`
	ObjectStorageUseProxy         *bool         `json:"objectStorageUseProxy,omitempty"`
	ObjectStorageSetPublicRead    *bool         `json:"objectStorageSetPublicRead,omitempty"`
	ObjectStorageS3ForcePathStyle *bool         `json:"objectStorageS3ForcePathStyle,omitempty"`
	DeeplAuthKey                  *string       `json:"deeplAuthKey,omitempty"`
	DeeplIsPro                    *bool         `json:"deeplIsPro,omitempty"`
}

type MetaAd struct {
	ID       string `json:"id"`
	URL      string `json:"url"`
	Place    string `json:"place"`
	Ratio    int    `json:"ratio"`
	ImageURL string `json:"imageUrl"`
}

type MetaEmoji struct {
	ID       string      `json:"id"`
	Aliases  []string    `json:"aliases"`
	Name     string      `json:"name"`
	Category null.String `json:"category"`
	Host     null.String `json:"host"`
	URL      string      `json:"url"`
}

type MetaFeatures struct {
	Registration           bool `json:"registration"`
	LocalTimeLine          bool `json:"localTimeLine"`
	GlobalTimeLine         bool `json:"globalTimeLine"`
	EmailRequiredForSignup bool `json:"emailRequiredForSignup"`
	Elasticsearch          bool `json:"elasticsearch"`
	HCaptcha               bool `json:"hcaptcha"`
	Recaptcha              bool `json:"recaptcha"`
	ObjectStorage          bool `json:"objectStorage"`
	Twitter                bool `json:"twitter"`
	Github                 bool `json:"github"`
	Discord                bool `json:"discord"`
	ServiceWorker          bool `json:"serviceWorker"`
	MiAuth                 bool `json:"miauth,omitempty"`
}
