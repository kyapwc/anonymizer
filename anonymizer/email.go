package anonymizer

import (
	"fmt"
	"regexp"
	"strings"
)

// Simple usage:
// ea := NewEmailAnonymizer("***")
// ea.Anonymize("a-a@a.b.c") -> ***@a.b.c

type emailAnonymizer struct {
	Replacement string
}

func NewEmailAnonymizer(replacement string) *emailAnonymizer {
	return &emailAnonymizer{
		Replacement: replacement,
	}
}

// simple method to check if the provided character is special
// (i.e: -/+/\\/\/)
func isSpecialCharacter(char byte) bool {
	specialCharacters := []byte{'-', '+', '/', '\\'} // Define a slice of special characters

	for _, specialChar := range specialCharacters {
		if char == specialChar {
			return true
		}
	}
	return false
}

func (ea *emailAnonymizer) Anonymize(s string) string {
	// initial email regex to check if there is valid email in provided string
	emailRegex := `[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`
	regex := regexp.MustCompile(emailRegex)

	// find all email addresses in the string
	emailMatches := regex.FindAllStringIndex(s, -1)

	// cumulative change to detect when should next occurrence of xxx@xxx.xxx in string
	cumulativeChange := 0
	// iterate over email matches and anonymize each one
	for _, match := range emailMatches {
		// extract the email address substring
		startIndex := match[0] + cumulativeChange
		endIndex := match[1] + cumulativeChange

		// special case if initial start of email has special character, we increase the startIndex + 1 to skip it
		if startIndex > 0 && isSpecialCharacter(s[startIndex]) {
			startIndex += 1
		}

		// grab the entire email
		email := s[startIndex:endIndex]
		// split the email and domain
		parts := strings.Split(email, "@")
		// grab the prefix
		partBeforeAlias := parts[0]

		// check if the prefix has a special character as the ending character before the alias symbol
		// if true, then we don't Anonymize it
		if isSpecialCharacter(partBeforeAlias[len(partBeforeAlias)-1]) {
			// break the loop
			continue
		}

		// just constant anonymized string
		anonymizedUsername := "..."

		// replace the original string with anonymizedUsername and succeeding parts of string
		s = s[:startIndex] + anonymizedUsername + "@" + parts[1] + s[endIndex:]

		// update the cumulative length change
		// have to - 1 as the index of string is different from len of string
		// len of string start from 1 and index of string start from 0
		cumulativeChange += len(anonymizedUsername) - 1
	}

	if len(emailMatches) == 0 {
		fmt.Println("No email matches for", s)
	}

	// if there are no matches for the provided string, just return the original string
	return s
}
