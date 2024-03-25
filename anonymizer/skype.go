package anonymizer

import (
	"fmt"
	"regexp"
	"strings"
)

// Simple usage:
// NewSkypeAnonymizer("#").Anonymize(`<a href="skype:loren?call"/>`) -> <a href="skype:#?call"/>

type skypeAnonymizer struct {
	Replacement string
}

func NewSkypeAnonymizer(replacement string) *skypeAnonymizer {
	return &skypeAnonymizer{
		Replacement: replacement,
	}
}

func (sa *skypeAnonymizer) Anonymize(s string) string {
	originalString := s
	usernameRegex := `skype:\w*`
	regex := regexp.MustCompile(usernameRegex)

	// find first occurence
	usernameMatches := regex.FindAllStringIndex(s, -1)
	fmt.Println("matches", usernameMatches)
	// Better call skype:MySkype or at least skype:YourAdmin

	if len(usernameMatches) > 0 {
		for _, match := range usernameMatches {
			startIndex := match[0]
			endIndex := match[1]

			username := originalString[startIndex:endIndex]

			parts := strings.Split(username, ":")

			s = strings.Replace(s, parts[1], "#", 1)
		}
	}

	return s
}
