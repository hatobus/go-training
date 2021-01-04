package detector

import (
	"errors"
	"strings"
)

var magicNumbers = map[string]string{
	"PK":    "zip",
	"ustar": "tar",
}

func DetectFileType(data []byte) (string, error) {
	for magic, mime := range magicNumbers {
		if strings.HasPrefix(string(data), magic) {
			return mime, nil
		}
	}
	return "", errors.New("invalid file type, this program supports zip or tar file")
}
