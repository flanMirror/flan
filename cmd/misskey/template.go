package main

import (
	"html/template"
	"log"
	"strings"

	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data"
	"random.chars.jp/git/misskey/db"
	"random.chars.jp/git/misskey/db/orm"
	"random.chars.jp/git/misskey/spec"
)

const (
	apiDocPath = "assets/public/static-assets/redoc.html"
)

var (
	apiDoc []byte

	biosTemplate  *data.TemplatePayload[map[string]string]
	cliTemplate   *data.TemplatePayload[map[string]string]
	flushTemplate *data.TemplatePayload[any]

	baseTemplate *data.TemplatePayload[map[string]string]
)

func init() {
	if d, err := assets.ReadFile(apiDocPath); err != nil {
		log.Printf("error caching asset %s: %s", apiDocPath, err)
	} else {
		apiDoc = d
	}
}

func templateSetup(templates *template.Template) {
	version := map[string]string{
		"version": spec.Target,
	}
	biosTemplate = data.NewTemplate[map[string]string](templates.Lookup("bios.tmpl"))
	cliTemplate = data.NewTemplate[map[string]string](templates.Lookup("cli.tmpl"))
	biosTemplate.Set(version)
	cliTemplate.Set(version)
	flushTemplate = data.NewTemplate[any](templates.Lookup("flush.tmpl"))
	flushTemplate.Set(nil)

	baseTemplate = data.NewTemplate[map[string]string](templates.Lookup("base.tmpl"))
}

func refreshMetaTemplates() {
	baseTemplate.Set(commonFormattingParameters.Get())
}

// commonFormattingParameters caches a map[string]string data structure
// containing common fields in base-derived templates
var commonFormattingParameters = data.New[map[string]string]()

const (
	instanceNameFallback = "Misskey"
	iconFallback         = "/favicon.ico"
	iconAppleFallback    = "/apple-touch-icon.png"
	descriptionFallback  = "✨🌎✨ A interplanetary communication platform ✨🚀✨"
	bannerFallback       = ""
)

func init() {
	db.Meta.Register(func(metum *orm.Metum) {
		//.instanceName - name of instance, must fall back to "Misskey" if invalid
		//.icon         - icon URL of instance, must fall back to "/favicon.ico" if invalid
		//.iconApple    - icon URL of instance, must fall back to "/apple-touch-icon.png" if invalid
		//.description  - instance description, must fall back to "✨🌎✨ A interplanetary communication platform ✨🚀✨" if invalid
		//.banner       - banner URL of instance, does not appear to have any form of fallback
		value := make(map[string]string, 5)
		if metum.Name.Valid {
			value["instanceName"] = metum.Name.String
		} else {
			value["instanceName"] = instanceNameFallback
		}
		if metum.IconUrl.Valid {
			value["icon"] = metum.IconUrl.String
			value["iconApple"] = metum.IconUrl.String
		} else {
			value["icon"] = iconFallback
			value["iconApple"] = iconAppleFallback
		}
		if metum.Description.Valid {
			value["description"] = metum.Description.String
		} else {
			value["description"] = descriptionFallback
		}
		if metum.BannerUrl.Valid {
			value["banner"] = metum.BannerUrl.String
		} else {
			value["banner"] = bannerFallback
		}
		commonFormattingParameters.Set(value)

		refreshMetaTemplates()
	})
}

func copyCommon(target map[string]any) {
	value := commonFormattingParameters.Get()

	target["instanceName"] =
		value["instanceName"]
	target["icon"] =
		value["icon"]
	target["iconApple"] =
		value["iconApple"]
	target["description"] =
		value["description"]
	target["banner"] =
		value["banner"]
}

func getUserParameters(user *orm.User, profile *orm.UserProfile, atRoot bool) map[string]any {
	//.atRoot                - whether the request is directed and the root of the user page
	//.userReference         - reference string of the user, "$DISPLAY_NAME (@$USERNAME)" or "@$USERNAME" if display name is invalid
	//.hasProfileDescription - whether profile.Description is valid
	//.profileDescription    - value of profile.Description
	//.url                   - value of config.URL
	//.noIndex               - true if user is a remote user or has profile.noCrawle set
	//.localUser             - true if user.Host is invalid
	//.userFull              - full user mention string without @ prefix
	//.userUsername          - value of user.Username
	//.userID                - value of user.ID
	//.userTwitterExists     - Deprecated, upstream forgot to remove it from templates
	//.userTwitterScreenName - Deprecated, upstream forgot to remove it from templates
	//.hasUserAvatarURL      - whether user.AvatarUrl is valid
	//.userAvatarURL         - value of user.AvatarUrl
	//.hasUserURI            - whether user.URI is valid
	//.userURI               - value of user.URI
	//.hasProfileURL         - whether profile.URL is valid
	//.profileURL            - value of profile.URL
	//.me                    - http(s) URLs in fields

	// 5 common and 18 specific, but 2 are deprecated
	value := make(map[string]interface{}, 21)
	copyCommon(value)

	value["atRoot"] = atRoot
	if user.Name.Valid {
		value["userReference"] = user.Name.String + " (@" + user.Username + ")"
	} else {
		value["userReference"] = "@" + user.Username
	}
	value["hasProfileDescription"] = profile.Description.Valid
	value["profileDescription"] = profile.Description.String
	value["url"] = config.URL

	userFull := user.Username
	value["noIndex"] = profile.NoCrawle
	if user.Host.Valid {
		userFull += "@" + user.Host.String
		value["noIndex"] = true
		value["localUser"] = false
	} else {
		value["localUser"] = true
	}
	value["userFull"] = userFull

	value["userUsername"] = user.Username
	value["userID"] = user.ID

	// twitter stuff appears to be deprecated, however upstream forgot to remove them from templates

	value["hasUserAvatarURL"] = user.AvatarUrl.Valid
	value["userAvatarURL"] = user.AvatarUrl.String
	value["hasUserURI"] = user.URI.Valid
	value["userURI"] = user.URI.String
	value["hasProfileURL"] = profile.URL.Valid
	value["profileURL"] = profile.URL.String

	var me []string
	var fields data.ProfileFields
	if err := profile.Fields.Unmarshal(&fields); err == nil {
		for _, field := range fields {
			if strings.HasPrefix(field.Value, "https:") ||
				strings.HasPrefix(field.Value, "http:") {
				me = append(me, field.Value)
			}
		}
	}
	value["me"] = me

	return value
}
