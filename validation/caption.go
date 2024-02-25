package validation

import "regexp"

// ValidateCaptionLength validates the length of the caption.
func ValidateCaptionLength(caption string) error {
	reHTML := regexp.MustCompile(`<[^>]+>`)
	caption = reHTML.ReplaceAllString(caption, "")
	reMD := regexp.MustCompile(`[_*{}\[\]()~` + "`" + `>|]+|\[.*?\]\(.*?\)|http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\\(\\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	caption = reMD.ReplaceAllString(caption, "")

	if len(caption) > 1024 {
		return ErrCaptionLength
	}
	return nil
}
