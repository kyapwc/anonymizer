package anonymizer

import (
	"errors"
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
	// TODO: Return text with anonymized phone number
	return ""
}
