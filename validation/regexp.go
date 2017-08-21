package validation

import "regexp"

const (
	regexEmailPattern       = `(?i)[A-Z0-9._%+-]+@(?:[A-Z0-9-]+\.)+[A-Z]{2,6}`
	regexStrictEmailPattern = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
		`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
		`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
		`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`
	regexURLPattern = `(ftp|http|https):\/\/(\w+:{0,1}\w*@)?(\S+)(:[0-9]+)?(\/|\/([\w#!:.?+=&%@!\-\/]))?`
)

var (
	regexEmail       *regexp.Regexp
	regexStrictEmail *regexp.Regexp
	regexURL         *regexp.Regexp
)

func init() {
	regexEmail = regexp.MustCompile(regexEmailPattern)
	regexStrictEmail = regexp.MustCompile(regexStrictEmailPattern)
	regexURL = regexp.MustCompile(regexURLPattern)
}

// IsEmail validates string is an email address, if not return false
// basically validation can match 99% cases
func IsEmail(email string) bool {
	return regexEmail.MatchString(email)
}

// IsEmailRFC validates string is an email address, if not return false
// this validation omits RFC 2822
func IsEmailRFC(email string) bool {
	return regexStrictEmail.MatchString(email)
}

// IsURL validates string is a url link, if not return false
// simple validation can match 99% cases
func IsURL(url string) bool {
	return regexURL.MatchString(url)
}
