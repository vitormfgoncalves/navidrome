package gravatar

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/navidrome/navidrome/utils/number"
)

const baseUrl = "https://www.gravatar.com/avatar"
const defaultSize = 80
const maxSize = 2048

func Url(email string, size int) string {
	email = strings.ToLower(email)
	email = strings.TrimSpace(email)
	hash := md5.Sum([]byte(email))
	if size < 1 {
		size = defaultSize
	}
	size = number.Min(maxSize, size)

	return fmt.Sprintf("%s/%x?s=%d", baseUrl, hash, size)
}
