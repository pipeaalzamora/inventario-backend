package shared

import "strings"

var MimeTypes = map[string]interface{}{
	"application/epub+zip":                            "epub",
	"application/gzip":                                "gz",
	"application/java-archive":                        "jar",
	"application/json":                                "json",
	"application/ld+json":                             "jsonld",
	"application/msword":                              "doc",
	"application/octet-stream":                        "bin",
	"application/ogg":                                 "ogx",
	"application/pdf":                                 "pdf",
	"application/rtf":                                 "rtf",
	"application/vnd.amazon.ebook":                    "azw",
	"application/vnd.apple.installer+xml":             "mpkg",
	"application/vnd.ms-excel":                        "xls",
	"application/vnd.ms-fontobject":                   "eot",
	"application/vnd.ms-powerpoint":                   "ppt",
	"application/vnd.oasis.opendocument.presentation": "odp",
	"application/vnd.oasis.opendocument.spreadsheet":  "ods",
	"application/vnd.oasis.opendocument.text":         "odt",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": "pptx",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         "xlsx",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   "docx",
	"application/vnd.rar":             "rar",
	"application/vnd.visio":           "vsd",
	"application/vnd.mozilla.xul+xml": "xul",
	"application/x-abiword":           "abw",
	"application/x-bzip":              "bz",
	"application/x-bzip2":             "bz2",
	"application/x-cdf":               "cda",
	"application/x-csh":               "csh",
	"application/x-freearc":           "arc",
	"application/x-httpd-php":         "php",
	"application/x-sh":                "sh",
	"application/x-tar":               "tar",
	"application/x-7z-compressed":     "7z",
	"audio/3gpp":                      "3gp",
	"audio/3gpp2":                     "3g2",
	"audio/aac":                       "aac",
	"audio/midi":                      "mid",
	"audio/mpeg":                      "mp3",
	"audio/ogg":                       "opus",
	"audio/wav":                       "wav",
	"audio/webm":                      "weba",
	"font/otf":                        "otf",
	"font/ttf":                        "ttf",
	"font/woff":                       "woff",
	"font/woff2":                      "woff2",
	"image/avif":                      "avif",
	"image/bmp":                       "bmp",
	"image/gif":                       "gif",
	"image/jpeg":                      []string{"jpeg", "jpg"},
	"image/png":                       "png",
	"image/svg+xml":                   "svg",
	"image/tiff":                      "tiff",
	"image/vnd.microsoft.icon":        "ico",
	"image/webp":                      "webp",
	"text/calendar":                   "ics",
	"text/css":                        "css",
	"text/csv":                        "csv",
	"text/html":                       "html",
	"text/javascript":                 "js",
	"text/plain":                      "txt",
	"video/3gpp":                      "3gp",
	"video/3gpp2":                     "3g2",
	"video/mp2t":                      "ts",
	"video/mp4":                       "mp4",
	"video/mpeg":                      "mpeg",
	"video/ogg":                       "ogv",
	"video/webm":                      "webm",
	"video/x-msvideo":                 "avi",
	"application/xhtml+xml":           "xhtml",
	"application/xml":                 "xml",
}

func GetMimeType(fileName string) string {
	ext := strings.ToLower(fileName[strings.LastIndex(fileName, "."):])
	extension := strings.Replace(ext, ".", "", 1)
	for key, value := range MimeTypes {
		switch v := value.(type) {
		case string:
			if v == extension {
				return key
			}
		case []string:
			for _, ext := range v {
				if ext == extension {
					return key
				}
			}
		}
	}

	return "application/octet-stream"
}
