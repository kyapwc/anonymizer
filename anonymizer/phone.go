package anonymizer

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Simple usage:
// pa := NewPhoneAnonymizer()
// pa.Anonymize("+48 666 777 888") -> +48 666 777 XXX
// pa.SetLastDigits(5)
// pa.Anonymize("+48 666 777 888") -> +48 666 7.. ...

type phoneAnonymizer struct {
	replacement rune
	lastDigits  int
}

func NewPhoneAnonymizer() *phoneAnonymizer {
	return &phoneAnonymizer{
		replacement: 'X',
		lastDigits:  3,
	}
}

func (pa *phoneAnonymizer) SetLastDigits(num int) error {
	if num < 0 || num > 9 {
		return errors.New("phoneAnonymizer: lastDigits must be >= 0 and <= 9")
	}

	pa.lastDigits = num
	return nil
}

func (pa *phoneAnonymizer) SetReplacement(newReplacement string) *phoneAnonymizer {
	pa.replacement = []rune(newReplacement)[0]
	return pa
}

func (pa *phoneAnonymizer) Anonymize(s string) string {
	// early return if the lastDigits == 0 to reduce on useless computation below
	if pa.lastDigits == 0 {
		return s
	}
	phoneRegex := `\+\d{2}\s\d{3}\s\d{3}\s\d{3}`
	regex := regexp.MustCompile(phoneRegex)

	// find all phone number matches
	matches := regex.FindAllStringIndex(s, -1)

	// instantiate a buffer to store the ananonymized string and the replaced string
	var buffer bytes.Buffer
	lastIndex := 0

	// Iterate over each phone number match
	for _, match := range matches {
		start := match[0]
		end := match[1]

		// appending the string before phone number to the current buffer
		// i.e: (Lorem) +48 bla bla bla bla
		buffer.WriteString(s[lastIndex:start])

		// extract phone number from original string provided
		phoneNumber := s[start:end]

		// dynamically define the regex pattern to match the provided lastDigits
		lastDigitsPattern := regexp.MustCompile(`\d{` + strconv.Itoa(pa.lastDigits) + `}$`)

		// replace all mention with provided `replacement` based on the lastDigitsPattern
		anonymizedPhoneNumber := lastDigitsPattern.ReplaceAllStringFunc(phoneNumber, func(match string) string {
			return strings.Repeat(string(pa.replacement), len(match))
		})

		// append the produced anonymizedPhoneNumber into the buffer
		buffer.WriteString(anonymizedPhoneNumber)

		// update the last index to trace current phone number update
		lastIndex = end
	}

	// append the suffix part of the original string back into buffer
	buffer.WriteString(s[lastIndex:])

	// return the entire stringified buffer as response
	return buffer.String()
}
