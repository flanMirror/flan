package drivefile

import (
	"strings"

	"github.com/volatiletech/null/v8"
	"random.chars.jp/git/misskey/api/response"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/db/orm"
	"random.chars.jp/git/misskey/prelude"
)

func ValidateFilename(name string) bool {
	return len(strings.TrimSpace(name)) > 0 &&
		len(name) <= 200 &&
		!strings.Contains(name, "..") &&
		!strings.ContainsAny(name, "/\\")
}

func GetPublicProperties(file *orm.DriveFile) (response.DriveFileProperties, error) {
	var properties response.DriveFileProperties

	if err := file.Properties.Unmarshal(&properties); err != nil {
		return response.DriveFileProperties{}, err
	}

	if properties.Orientation != nil {
		if *properties.Orientation >= 5 {
			properties.Width, properties.Height = properties.Height, properties.Width
		}
		properties.Orientation = nil
	}

	return properties, nil
}

var imageMime = []string{"image/png", "image/apng", "image/gif", "image/jpeg", "image/webp", "image/svg+xml"}

// GetPublicURL gets public URL of a drive file, thumbnail false by default, meta nil by default
func GetPublicURL(file *orm.DriveFile, thumbnail bool, meta *orm.Metum) string {
	if file.URI.Valid {
		if file.UserHost.Valid && config.System.MediaProxy != "" {
			q := map[string]string{
				"url": file.URI.String,
			}
			if thumbnail {
				q["thumbnail"] = "1"
			}
			return prelude.AppendQuery(config.System.MediaProxy, prelude.Query(q))
		}

		if file.IsLink && meta != nil && meta.ProxyRemoteFiles {
			var key null.String
			if thumbnail {
				key = file.ThumbnailAccessKey
			} else {
				key = file.WebpublicAccessKey
			}

			if key.Valid && !strings.ContainsRune(key.String, '/') {
				return config.URL + "/files/" + key.String
			}
		}
	}

	isImage := false
	for _, m := range imageMime {
		if file.Type == m {
			isImage = true
			break
		}
	}

	if thumbnail {
		if file.ThumbnailUrl.Valid {
			return file.ThumbnailUrl.String
		}
		if !isImage {
			return ""
		}
	}
	if file.WebpublicUrl.Valid {
		return file.WebpublicUrl.String
	}
	return file.URL
}
