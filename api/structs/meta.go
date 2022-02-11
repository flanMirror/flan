package structs

import "github.com/volatiletech/null/v8"

type Meta struct {
	MaintainerName               null.String   `json:"maintainerName"`
	MaintainerEmail              null.String   `json:"maintainerEmail"`
	Version                      string        `json:"version"`
	Name                         string        `json:"name"`
	URI                          string        `json:"uri"`
	Description                  null.String   `json:"description"`
	Langs                        []string      `json:"langs"`
	TosURL                       null.String   `json:"tosUrl"`
	RepositoryURL                string        `json:"repositoryUrl"`
	FeedbackURL                  string        `json:"feedbackUrl"`
	Secure                       bool          `json:"secure"`
	DisableRegistration          bool          `json:"disableRegistration"`
	DisableLocalTimeline         bool          `json:"disableLocalTimeline"`
	DisableGlobalTimeline        bool          `json:"disableGlobalTimeline"`
	DriveCapacityPerLocalUserMb  int           `json:"driveCapacityPerLocalUserMb"`
	DriveCapacityPerRemoteUserMb int           `json:"driveCapacityPerRemoteUserMb"`
	CacheRemoteFiles             *bool         `json:"cacheRemoteFiles,omitempty"`
	ProxyRemoteFiles             *bool         `json:"proxyRemoteFiles,omitempty"`
	EmailRequiredForSignup       bool          `json:"emailRequiredForSignup"`
	EnableHcaptcha               bool          `json:"enableHcaptcha"`
	HcaptchaSiteKey              null.String   `json:"hcaptchaSiteKey"`
	EnableRecaptcha              bool          `json:"enableRecaptcha"`
	RecaptchaSiteKey             null.String   `json:"recaptchaSiteKey"`
	SwPublickey                  null.String   `json:"swPublickey"`
	MascotImageURL               string        `json:"mascotImageUrl"`
	BannerURL                    string        `json:"bannerUrl"`
	ErrorImageURL                string        `json:"errorImageUrl"`
	IconURL                      null.String   `json:"iconUrl"`
	BackgroundImageURL           null.String   `json:"backgroundImageUrl"`
	LogoImageURL                 null.String   `json:"logoImageUrl"`
	MaxNoteTextLength            int           `json:"maxNoteTextLength"`
	Emojis                       []MetaEmoji   `json:"emojis"`
	Ads                          []MetaAd      `json:"ads"`
	RequireSetup                 *bool         `json:"requireSetup,omitempty"`
	EnableEmail                  bool          `json:"enableEmail"`
	EnableTwitterIntegration     bool          `json:"enableTwitterIntegration"`
	EnableGithubIntegration      bool          `json:"enableGithubIntegration"`
	EnableDiscordIntegration     bool          `json:"enableDiscordIntegration"`
	EnableServiceWorker          bool          `json:"enableServiceWorker"`
	TranslatorAvailable          bool          `json:"translatorAvailable"`
	ProxyAccountName             *null.String  `json:"proxyAccountName,omitempty"`
	Features                     *MetaFeatures `json:"features,omitempty"`
	PinnedPages                  []string      `json:"pinnedPages,omitempty"`
	PinnedClipID                 *null.String  `json:"pinnedClipId,omitempty"`
}

type MetaAdmin struct {
	UseStarForReactionFallback    bool        `json:"useStarForReactionFallback"`
	PinnedUsers                   []string    `json:"pinnedUsers"`
	HiddenTags                    []string    `json:"hiddenTags"`
	BlockedHosts                  []string    `json:"blockedHosts"`
	HCaptchaSecretKey             null.String `json:"hcaptchaSecretKey"`
	RecaptchaSecretKey            null.String `json:"recaptchaSecretKey"`
	ProxyAccountID                null.String `json:"proxyAccountId"`
	TwitterConsumerKey            null.String `json:"twitterConsumerKey"`
	TwitterConsumerSecret         null.String `json:"twitterConsumerSecret"`
	GithubClientID                null.String `json:"githubClientId"`
	GithubClientSecret            null.String `json:"githubClientSecret"`
	DiscordClientID               null.String `json:"discordClientId"`
	DiscordClientSecret           null.String `json:"discordClientSecret"`
	SummalyProxy                  null.String `json:"summalyProxy"`
	Email                         null.String `json:"email"`
	SmtpSecure                    bool        `json:"smtpSecure"`
	SmtpHost                      null.String `json:"smtpHost"`
	SmtpPort                      null.Int    `json:"smtpPort"`
	SmtpUser                      null.String `json:"smtpUser"`
	SmtpPass                      null.String `json:"smtpPass"`
	SwPrivateKey                  null.String `json:"swPrivateKey"`
	UseObjectStorage              bool        `json:"useObjectStorage,omitempty"`
	ObjectStorageBaseURL          null.String `json:"objectStorageBaseUrl"`
	ObjectStorageBucket           null.String `json:"objectStorageBucket"`
	ObjectStoragePrefix           null.String `json:"objectStoragePrefix"`
	ObjectStorageEndpoint         null.String `json:"objectStorageEndpoint"`
	ObjectStorageRegion           null.String `json:"objectStorageRegion"`
	ObjectStoragePort             null.Int    `json:"objectStoragePort"`
	ObjectStorageAccessKey        null.String `json:"objectStorageAccessKey"`
	ObjectStorageSecretKey        null.String `json:"objectStorageSecretKey"`
	ObjectStorageUseSSL           bool        `json:"objectStorageUseSSL"`
	ObjectStorageUseProxy         bool        `json:"objectStorageUseProxy"`
	ObjectStorageSetPublicRead    bool        `json:"objectStorageSetPublicRead"`
	ObjectStorageS3ForcePathStyle bool        `json:"objectStorageS3ForcePathStyle"`
	DeeplAuthKey                  null.String `json:"deeplAuthKey"`
	DeeplIsPro                    bool        `json:"deeplIsPro"`

	Meta
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
